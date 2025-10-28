
#include <iostream>
#include <vector>
#include <sys/ioctl.h>
#include <unistd.h>
#include <random>
#include <chrono>
#include <thread>

using namespace std;

// Render the grid in the terminal
void render(const vector<vector<int>>& grid) {
    cout << "\033[H\033[J"; // Clear screen
    for (const auto& row : grid) {
        for (int val : row) {
            cout << (val ? "■" : " "); // alive=■, dead=space
        }
        cout << "\n";
    }
}

// Count live neighbors
int countNeighbors(const vector<vector<int>>& grid, int r, int c) {
    int count = 0;
    for (int dr = -1; dr <= 1; dr++) {
        for (int dc = -1; dc <= 1; dc++) {
            if (dr == 0 && dc == 0) continue; // skip self
            int nr = r + dr;
            int nc = c + dc;
            if (nr >= 0 && nr < grid.size() && nc >= 0 && nc < grid[0].size()) {
                count += grid[nr][nc];
            }
        }
    }
    return count;
}

// Update the grid according to Game of Life rules
vector<vector<int>> update(const vector<vector<int>>& grid) {
    vector<vector<int>> newGrid = grid;
    for (int r = 0; r < grid.size(); r++) {
        for (int c = 0; c < grid[0].size(); c++) {
            int neighbors = countNeighbors(grid, r, c);
            if (grid[r][c] == 1) {
                if (neighbors < 2 || neighbors > 3) newGrid[r][c] = 0; // dies
            } else {
                if (neighbors == 3) newGrid[r][c] = 1; // becomes alive
            }
        }
    }
    return newGrid;
}

int main() {
    // Get terminal size
    struct winsize w;
    if (ioctl(STDOUT_FILENO, TIOCGWINSZ, &w) == -1) {
        perror("ioctl");
        return 1;
    }

    int rows = w.ws_row - 1; // leave one line for iteration count
    int cols = w.ws_col;

    // Initialize grid
    vector<vector<int>> grid(rows, vector<int>(cols, 0));

    // Randomly fill ~20% of cells
    random_device rd;
    mt19937 gen(rd());
    uniform_int_distribution<> rowDist(0, rows - 1);
    uniform_int_distribution<> colDist(0, cols - 1);

    int numCells = rows * cols / 10; // ~20%
    for (int i = 0; i < numCells; i++) {
        int r = rowDist(gen);
        int c = colDist(gen);
        grid[r][c] = 1;
    }

    // Main loop
    int iterations = 10000;
    for (int i = 0; i < iterations; i++) {
        render(grid);
        //cout << "Iteration: " << i + 1 << "/" << iterations << "\n";
        grid = update(grid);
        //this_thread::sleep_for(chrono::milliseconds(0)); // 50ms delay
    }

    return 0;
}

