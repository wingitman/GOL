#include <stdio.h>
#include <unistd.h>

#define WIDTH 40
#define HEIGHT 20

int grid[HEIGHT][WIDTH];
int next[HEIGHT][WIDTH];

int count_neighbors(int g[HEIGHT][WIDTH], int y, int x) {
  int count = 0;
  for (int dy = -1; dy <= 1; dy++) {
    for (int dx = -1; dx <= 1; dx++) {
      if (dy == 0 && dx == 0) continue;

      int ny = y + dy;
      int nx = x + dx;

      if (ny >= 0 && ny < HEIGHT && nx >= 0 && nx < WIDTH) {
        count += g[ny][nx];
      }
    }
  }
  return count;
}

void step(int g[HEIGHT][WIDTH], int n[HEIGHT][WIDTH]) {
  for (int y = 0; y < HEIGHT; y++) {
    for (int x = 0; x < WIDTH; x++) {
      int neighbors = count_neighbors(g, y, x);

      if (g[y][x] == 1) {
        n[y][x] = (neighbors == 2 || neighbors == 3);
      } else {
        n[y][x] = neighbors == 3;
      }
    }
  }
}

void print_grid(int g[HEIGHT][WIDTH]) {
  for (int y = 0; y < HEIGHT; y++) {
    for (int x = 0; x < WIDTH; x++) {
      printf(g[y][x] ? "█" : " ");
    }
    printf("\n");
  }
}

int main() {
  int grid[HEIGHT][WIDTH] = {0};
  grid[1][2] = grid[2][3] = grid[3][1] = grid[3][2] = grid[3][3] = 1;

  int next[HEIGHT][WIDTH];
  while (1) {
    print_grid(grid);
    step(grid, next);

    for (int y = 0; y < HEIGHT; y++) {
      for (int x = 0; x < WIDTH; x++) {
        grid[y][x] = next[y][x];
      }
    }
    usleep(100000);
    printf("\033[H\033[J");
  }
  return 0;
}
