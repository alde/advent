const std = @import("std");
const shared = @import("../shared.zig");

pub fn part1() !void {
    var logger = try shared.initLogger();
    defer logger.deinit();
    logger.info("starting", .{ .day = 1, .part = 1 });
    const inputData = try shared.readInput("src/day1/input.txt");
    const result: i32 = countUpDown(inputData);
    logger.info("completed", .{ .day = 1, .part = 1, .result = result });
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
    var logger = try shared.initLogger();
    defer logger.deinit();
    logger.info("starting", .{ .day = 1, .part = 2 });
    const inputData = try shared.readInput("src/day1/input.txt");
    const result: i32 = findBasement(inputData);
    logger.info("completed", .{ .day = 1, .part = 2, .result = result });
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
        try testing.expectEqual(0, countUpDown("(())"));
        try testing.expectEqual(0, countUpDown("()()"));
        try testing.expectEqual(1, countUpDown("(()"));
        try testing.expectEqual(-1, countUpDown("())"));
        try testing.expectEqual(3, countUpDown("((("));
        try testing.expectEqual(3, countUpDown("(()(()("));
        try testing.expectEqual(3, countUpDown("))((((("));
    }

    {
        try testing.expectEqual(1, findBasement(")"));
        try testing.expectEqual(5, findBasement("()())"));
    }
}
