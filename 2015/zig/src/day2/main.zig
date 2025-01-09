const std = @import("std");
const shared = @import("../shared.zig");

pub fn part1() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 2 Part 1\n", .{});

    var line_it = try shared.iterLines("src/day2/input.txt");
    defer line_it.deinit();

    var result: i32 = 0;
    while (try line_it.next()) |line| {
        const dim = try parseDimensions(line);
        result += neededPaper(dim[0], dim[1], dim[2]);
    }

    try stdout.print("Result: {}\n", .{result});
}

fn neededPaper(length: i32, width: i32, height: i32) i32 {
    const a = 2 * length * width;
    const b = 2 * width * height;
    const c = 2 * height * length;

    const smallestSide: i32 = min(min(a, b), c) >> 1;

    return a + b + c + smallestSide;
}
fn min(a: i32, b: i32) i32 {
    return if (a < b) a else b;
}

fn parseDimensions(input: []const u8) ![3]i32 {
    var it = std.mem.splitScalar(u8, input, 'x');
    var mem: [3]i32 = [3]i32{ 0, 0, 0 };
    var curr: usize = 0;
    while (it.next()) |piece| {
        mem[curr] = try std.fmt.parseInt(i32, piece, 10);
        curr += 1;
    }
    return mem;
}

pub fn part2() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 2 Part 2\n", .{});

    var line_it = try shared.iterLines("src/day2/input.txt");
    defer line_it.deinit();

    var result: i32 = 0;
    while (try line_it.next()) |line| {
        const dim = try parseDimensions(line);
        result += ribbonLength(dim[0], dim[1], dim[2]);
    }

    try stdout.print("Result: {}\n", .{result});
}

fn ribbonLength(a: i32, b: i32, c: i32) i32 {
    var smallestFace = a + a + b + b;
    if (b + b + c + c < smallestFace) {
        smallestFace = b + b + c + c;
    }
    if (a + a + c + c < smallestFace) {
        smallestFace = a + a + c + c;
    }

    const extra = a * b * c;

    return smallestFace + extra;
}

test "day 2" {
    const testing = std.testing;

    {
        try testing.expectEqual(58, neededPaper(2, 3, 4));
        try testing.expectEqual(43, neededPaper(1, 1, 10));

        try testing.expectEqual([3]i32{ 2, 3, 4 }, parseDimensions("2x3x4"));
        try testing.expectEqual([3]i32{ 1, 1, 10 }, parseDimensions("1x1x10"));
    }

    {
        try testing.expectEqual(34, ribbonLength(2, 3, 4));
        try testing.expectEqual(14, ribbonLength(1, 1, 10));
    }
}
