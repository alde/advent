const std = @import("std");
const shared = @import("../shared.zig");
const zlog = @import("zlog");

const Log = zlog.Logger(.info, .{});
const ColorHandler = zlog.ColorHandler.Handler(.{ .timestamp = .rfc3339 });

pub fn part1() !void {
    var handler = ColorHandler.init(zlog.stderr);
    var logger = try Log.init(.{ .handler = handler.handler(), .allocator = std.heap.page_allocator });
    defer logger.deinit();
    logger.info("starting", .{ .day = 3, .part = 1 });
    const inputData = try shared.readInput("src/day3/input.txt");
    const result: usize = try countHouses(inputData);

    logger.info("completed", .{ .day = 3, .part = 1, .result = result });
}

fn countHouses(directions: []const u8) !usize {
    const allocator = std.heap.page_allocator;
    const Point = struct { x: i32, y: i32 };
    var map = std.AutoHashMap(Point, u32).init(allocator);
    defer map.deinit();

    var pos: Point = .{ .x = 0, .y = 0 };
    try map.put(pos, 1);
    for (directions) |d| {
        switch (d) {
            'v' => pos.y = pos.y - 1,
            '^' => pos.y = pos.y + 1,
            '>' => pos.x = pos.x + 1,
            '<' => pos.x = pos.x - 1,
            else => continue,
        }
        const curr = try map.getOrPut(pos);
        if (curr.found_existing) {
            curr.value_ptr.* += 1;
        } else {
            curr.value_ptr.* = 1;
        }
    }

    return map.count();
}

pub fn part2() !void {
    var handler = ColorHandler.init(zlog.stderr);
    var logger = try Log.init(.{ .handler = handler.handler(), .allocator = std.heap.page_allocator });
    defer logger.deinit();
    logger.info("starting", .{ .day = 3, .part = 2 });
    const inputData = try shared.readInput("src/day3/input.txt");
    const result: usize = try countHousesWithRobot(inputData);

    logger.info("completed", .{ .day = 3, .part = 2, .result = result });
}

fn countHousesWithRobot(directions: []const u8) !usize {
    const allocator = std.heap.page_allocator;
    const Point = struct { x: i32, y: i32 };
    var map = std.AutoHashMap(Point, u32).init(allocator);
    defer map.deinit();

    var posSanta: Point = .{ .x = 0, .y = 0 };
    var posRobot: Point = .{ .x = 0, .y = 0 };
    try map.put(posSanta, 1);
    for (directions, 0..) |d, i| {
        if (i % 2 == 0) {
            switch (d) {
                'v' => posSanta.y = posSanta.y - 1,
                '^' => posSanta.y = posSanta.y + 1,
                '>' => posSanta.x = posSanta.x - 1,
                '<' => posSanta.x = posSanta.x + 1,
                else => continue,
            }
        } else {
            switch (d) {
                'v' => posRobot.y = posRobot.y - 1,
                '^' => posRobot.y = posRobot.y + 1,
                '>' => posRobot.x = posRobot.x - 1,
                '<' => posRobot.x = posRobot.x + 1,
                else => continue,
            }
        }

        const curr = if (i % 2 == 0) try map.getOrPut(posSanta) else try map.getOrPut(posRobot);
        if (curr.found_existing) {
            curr.value_ptr.* += 1;
        } else {
            curr.value_ptr.* = 1;
        }
    }

    return map.count();
}

test "day 3" {
    const testing = std.testing;

    {
        try testing.expectEqual(2, countHouses("v^"));
        try testing.expectEqual(2, countHouses(">"));
        try testing.expectEqual(4, countHouses("^>v<"));
        try testing.expectEqual(2, countHouses("^v^v^v^v^v"));
    }

    {
        try testing.expectEqual(3, countHousesWithRobot("v^"));
        try testing.expectEqual(3, countHousesWithRobot("^>v<"));
        try testing.expectEqual(11, countHousesWithRobot("^v^v^v^v^v"));
    }
}
