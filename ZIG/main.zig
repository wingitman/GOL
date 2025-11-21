const std = @import("std");

const GridSize = 20;
const Iterations = 100;

const Alive: u8 = 1;
const Dead: u8 = 0;

const Grid = [GridSize][GridSize]u8;

fn printGrid(grid: *const Grid, writer: anytype) !void {
    for (grid) |row| {
        for (row) |cell| {
            if (cell == Alive) {
                try writer.print("■", .{});
            } else {
                try writer.print("□", .{});
            }
        }
        try writer.print("\n", .{});
    }
    try writer.print("\n", .{});
}

fn countNeighbors(grid: *const Grid, x: usize, y: usize) u8 {
    var count: u8 = 0;

    var dx: i32 = -1;
    while (dx <= 1) : (dx += 1) {
        var dy: i32 = -1;
        while (dy <= 1) : (dy += 1) {
            if (dx == 0 and dy == 0) continue;

            const nx_i32 = @as(i32, @intCast(x)) + dx;
            const ny_i32 = @as(i32, @intCast(y)) + dy;

            if (nx_i32 >= 0 and ny_i32 >= 0 and
                nx_i32 < GridSize and ny_i32 < GridSize)
            {
                const nx = @as(usize, @intCast(nx_i32));
                const ny = @as(usize, @intCast(ny_i32));
                if (grid[nx][ny] == Alive)
                    count += 1;
            }
        }
    }

    return count;
}

fn nextGeneration(grid: *Grid) void {
    var newGrid: Grid = undefined;

    for (grid, 0..) |row, x| {
        for (row, 0..) |cell, y| {
            const neighbors = countNeighbors(grid, x, y);

            if (cell == Alive) {
                newGrid[x][y] = if (neighbors < 2 or neighbors > 3) Dead else Alive;
            } else {
                newGrid[x][y] = if (neighbors == 3) Alive else Dead;
            }
        }
    }

    grid.* = newGrid;
}

pub fn main() !void {
    const stdoutFile = std.fs.File.stdout();
    var buffer: [1024]u8 = undefined;
    var stdoutWriter = stdoutFile.writer(&buffer);
    const writer = &stdoutWriter.interface;

    var grid: Grid = [_][GridSize]u8{[_]u8{0} ** GridSize} ** GridSize;

    grid[10][10] = Alive;
    grid[11][11] = Alive;
    grid[12][9] = Alive;
    grid[12][10] = Alive;
    grid[12][11] = Alive;

    var i: usize = 0;
    while (i < Iterations) : (i += 1) {
        try writer.print("Generation {}:\n", .{i});
        try printGrid(&grid, writer);
        nextGeneration(&grid);
        std.time.sleep(200 * std.time.millisecond);
    }
}
