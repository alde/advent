const std = @import("std");
const shared = @import("../shared.zig");
const stdout = std.io.getStdOut().writer();
const testing = std.testing;

pub fn part1() !void {
    try stdout.print("Day 6 Part 1\n", .{});
    var line_it = try shared.iterLines("src/day6/input.txt");
    defer line_it.deinit();
    const n: usize = 1000;

    // Need to create the matrix on the heap because it's too big for the stack.
    // This is not apparent because `zig build` doesn't say anything - it just
    // crashes with a exit code 1 and no stack trace.
    const allocator = std.heap.page_allocator;

    const temp = try allocator.alloc([n][n]usize, 1);
    defer allocator.free(temp);

    // Initialize the slice
    for (&temp[0]) |*row| {
        for (row) |*val| {
            val.* = 0;
        }
    }
    var matrix: [n][n]usize = temp.ptr[0];

    while (try line_it.next()) |line| {
        const is = try processInstruction(line);
        // stdout.print("is.action {}\n", .{is.action}) catch unreachable;
        updateMatrix(n, n, &matrix, is);
    }
    const result = countLights(n, n, matrix);

    try stdout.print("Result: {}\n", .{result});
}

pub fn part2() !void {
    try stdout.print("Day 6 Part 2\n", .{});
    var line_it = try shared.iterLines("src/day6/input.txt");
    defer line_it.deinit();
    const n: usize = 1000;

    // Need to create the matrix on the heap because it's too big for the stack.
    // This is not apparent because `zig build` doesn't say anything - it just
    // crashes with a exit code 1 and no stack trace.
    const allocator = std.heap.page_allocator;

    const temp = try allocator.alloc([n][n]usize, 1);
    defer allocator.free(temp);

    // Initialize the slice
    for (&temp[0]) |*row| {
        for (row) |*val| {
            val.* = 0;
        }
    }
    var matrix: [n][n]usize = temp.ptr[0];

    while (try line_it.next()) |line| {
        const is = try processInstruction(line);
        // stdout.print("is.action {}\n", .{is.action}) catch unreachable;
        updateMatrix2(n, n, &matrix, is);
    }
    const result = countLights(n, n, matrix);

    try stdout.print("Result: {}\n", .{result});
}

const Action = enum { TURN_ON, TOGGLE, TURN_OFF };
const Point = struct {
    x: usize,
    y: usize,
};

const InstructionSet = struct {
    action: Action,
    start: Point,
    end: Point,
};

fn createPoint(s: []const u8) Point {
    var it = std.mem.split(u8, s, ",");
    var ints: [2]usize = [2]usize{ 0, 0 };
    var i: usize = 0;
    while (it.next()) |item| : (i += 1) {
        ints[i] = std.fmt.parseInt(usize, item, 10) catch unreachable;
    }
    return .{ .x = ints[0], .y = ints[1] };
}

fn processInstruction(instruction: []const u8) !InstructionSet {
    if (std.mem.eql(u8, instruction[0..7], "turn on")) {
        var it = std.mem.split(u8, instruction[8..], " ");
        var i: i32 = 0;
        var start: Point = .{ .x = 0, .y = 0 };
        var end: Point = .{ .x = 0, .y = 0 };
        while (it.next()) |item| : (i += 1) {
            if (i == 0) {
                start = createPoint(item);
            } else if (i == 2) {
                end = createPoint(item);
            }
        }
        return .{
            .action = Action.TURN_ON,
            .start = start,
            .end = end,
        };
    }
    if (std.mem.eql(u8, instruction[0..8], "turn off")) {
        var it = std.mem.split(u8, instruction[9..], " ");
        var i: i32 = 0;
        var start: Point = .{ .x = 0, .y = 0 };
        var end: Point = .{ .x = 0, .y = 0 };
        while (it.next()) |item| : (i += 1) {
            if (i == 0) {
                start = createPoint(item);
            } else if (i == 2) {
                end = createPoint(item);
            }
        }
        return .{
            .action = Action.TURN_OFF,
            .start = start,
            .end = end,
        };
    }
    if (std.mem.eql(u8, instruction[0..6], "toggle")) {
        var it = std.mem.split(u8, instruction[7..], " ");
        var i: i32 = 0;
        var start: Point = .{ .x = 0, .y = 0 };
        var end: Point = .{ .x = 0, .y = 0 };
        while (it.next()) |item| : (i += 1) {
            if (i == 0) {
                start = createPoint(item);
            } else if (i == 2) {
                end = createPoint(item);
            }
        }
        return .{
            .action = Action.TOGGLE,
            .start = start,
            .end = end,
        };
    }
    return error.Invalid;
}

fn updateMatrix(comptime X: usize, comptime Y: usize, mat: *[X][Y]usize, is: InstructionSet) void {
    for (is.start.x..(is.end.x + 1)) |x| {
        if (x == X) {
            break;
        }
        for (is.start.y..(is.end.y + 1)) |y| {
            if (y == Y) {
                break;
            }
            switch (is.action) {
                Action.TURN_ON => mat[x][y] = 1,
                Action.TURN_OFF => mat[x][y] = 0,
                Action.TOGGLE => {
                    if (mat[x][y] == 0) mat[x][y] = 1 else mat[x][y] = 0;
                },
            }
        }
    }
}

fn updateMatrix2(comptime X: usize, comptime Y: usize, mat: *[X][Y]usize, is: InstructionSet) void {
    for (is.start.x..(is.end.x + 1)) |x| {
        for (is.start.y..(is.end.y + 1)) |y| {
            switch (is.action) {
                Action.TURN_ON => mat[x][y] += 1,
                Action.TURN_OFF => {
                    if (mat[x][y] > 0) mat[x][y] -= 1 else mat[x][y] = 0;
                },
                Action.TOGGLE => mat[x][y] += 2,
            }
        }
    }
}

fn countLights(comptime X: usize, comptime Y: usize, mat: [X][Y]usize) usize {
    var count: usize = 0;
    for (mat) |row| {
        for (row) |light| {
            count += light;
        }
    }
    return count;
}

test "processInstruction" {
    try testing.expectEqual(InstructionSet{ .action = Action.TURN_ON, .start = Point{ .x = 0, .y = 0 }, .end = Point{ .x = 3, .y = 3 } }, processInstruction("turn on 0,0 through 3,3"));
    try testing.expectEqual(InstructionSet{ .action = Action.TURN_ON, .start = Point{ .x = 1, .y = 2 }, .end = Point{ .x = 99, .y = 3 } }, processInstruction("turn on 1,2 through 99,3"));
    try testing.expectEqual(InstructionSet{ .action = Action.TURN_OFF, .start = Point{ .x = 1, .y = 2 }, .end = Point{ .x = 99, .y = 3 } }, processInstruction("turn off 1,2 through 99,3"));
    try testing.expectEqual(InstructionSet{ .action = Action.TOGGLE, .start = Point{ .x = 1, .y = 2 }, .end = Point{ .x = 99, .y = 3 } }, processInstruction("toggle 1,2 through 99,3"));
}

test "updateMatrix" {
    {
        var mat = [3][3]usize{ .{ 0, 0, 0 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const expected = [3][3]usize{ .{ 1, 1, 1 }, .{ 1, 1, 1 }, .{ 1, 1, 1 } };
        const is: InstructionSet = InstructionSet{ .action = Action.TURN_ON, .start = Point{ .x = 0, .y = 0 }, .end = Point{ .x = 2, .y = 2 } };
        updateMatrix(3, 3, &mat, is);
        try testing.expectEqual(
            expected,
            mat,
        );
    }
    {
        var mat = [3][3]usize{ .{ 0, 0, 0 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const expected = [3][3]usize{ .{ 1, 1, 1 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const is: InstructionSet = InstructionSet{ .action = Action.TOGGLE, .start = Point{ .x = 0, .y = 0 }, .end = Point{ .x = 0, .y = 2 } };
        updateMatrix(3, 3, &mat, is);
        try testing.expectEqual(
            expected,
            mat,
        );
    }
}

test "updateMatrix2" {
    {
        var mat = [3][3]usize{ .{ 0, 0, 0 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const expected = [3][3]usize{ .{ 1, 1, 1 }, .{ 1, 1, 1 }, .{ 1, 1, 1 } };
        const is: InstructionSet = InstructionSet{ .action = Action.TURN_ON, .start = Point{ .x = 0, .y = 0 }, .end = Point{ .x = 2, .y = 2 } };
        updateMatrix2(3, 3, &mat, is);
        try testing.expectEqual(
            expected,
            mat,
        );
    }
    {
        var mat = [3][3]usize{ .{ 0, 0, 0 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const expected = [3][3]usize{ .{ 2, 2, 2 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const is: InstructionSet = InstructionSet{ .action = Action.TOGGLE, .start = Point{ .x = 0, .y = 0 }, .end = Point{ .x = 0, .y = 2 } };
        updateMatrix2(3, 3, &mat, is);
        try testing.expectEqual(
            expected,
            mat,
        );
    }
    {
        var mat = [3][3]usize{ .{ 2, 2, 2 }, .{ 1, 1, 1 }, .{ 0, 0, 0 } };
        const expected = [3][3]usize{ .{ 1, 1, 1 }, .{ 0, 0, 0 }, .{ 0, 0, 0 } };
        const is: InstructionSet = InstructionSet{ .action = Action.TURN_OFF, .start = Point{ .x = 0, .y = 0 }, .end = Point{ .x = 2, .y = 2 } };
        updateMatrix2(3, 3, &mat, is);
        try testing.expectEqual(
            expected,
            mat,
        );
    }
}

test "countLights" {
    const mat = [3][3]usize{ .{ 1, 1, 1 }, .{ 1, 1, 1 }, .{ 1, 1, 1 } };
    try testing.expectEqual(9, countLights(3, 3, mat));

    const mat2 = [3][3]usize{ .{ 1, 1, 1 }, .{ 1, 0, 1 }, .{ 1, 1, 1 } };
    try testing.expectEqual(8, countLights(3, 3, mat2));

    const mat3 = [3][3]usize{ .{ 0, 0, 0 }, .{ 1, 0, 1 }, .{ 0, 0, 0 } };
    try testing.expectEqual(2, countLights(3, 3, mat3));

    const mat4 = [3][3]usize{ .{ 2, 2, 2 }, .{ 1, 0, 1 }, .{ 0, 0, 0 } };
    try testing.expectEqual(8, countLights(3, 3, mat4));
}
