const std = @import("std");
const shared = @import("../shared.zig");

pub fn part1() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 1 Part 1\n", .{});
    const inputData = try shared.readInput("src/day1/input.txt");
    const result: i32 = countUpDown(inputData);
    try stdout.print("Result: {}\n", .{result});
}

fn countUpDown(line: []const u8) i32 {
    var curr: i32 = 0;
    for (line) |value| {
        if (value == '(') {
            curr += 1;
        }
        if (value == ')') {
            curr -= 1;
        }
    }

    return curr;
}

pub fn part2() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 1 Part 2\n", .{});
    const inputData = try shared.readInput("src/day1/input.txt");
    const result: i32 = findBasement(inputData);
    try stdout.print("Result: {}\n", .{result});
}

fn findBasement(line: []const u8) i32 {
    var curr: i32 = 0;
    var pos: i32 = 0;
    for (line) |value| {
        pos += 1;
        if (value == '(') {
            curr += 1;
        }
        if (value == ')') {
            curr -= 1;
        }
        if (curr < 0) {
            return pos;
        }
    }

    return curr;
}
test "day 1" {
    const testing = std.testing;

    {
        // countUpDown
        try testing.expectEqual(0, countUpDown("(())"));
        try testing.expectEqual(0, countUpDown("()()"));
        try testing.expectEqual(1, countUpDown("(()"));
        try testing.expectEqual(-1, countUpDown("())"));
        try testing.expectEqual(3, countUpDown("((("));
        try testing.expectEqual(3, countUpDown("(()(()("));
        try testing.expectEqual(3, countUpDown("))((((("));
    }

    {
        // findBasement
        try testing.expectEqual(1, findBasement(")"));
        try testing.expectEqual(5, findBasement("()())"));
    }
}
