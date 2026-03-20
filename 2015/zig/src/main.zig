const std = @import("std");
const zlog = @import("zlog");
const day1 = @import("day1/main.zig");
const day2 = @import("day2/main.zig");
const day3 = @import("day3/main.zig");
const day4 = @import("day4/main.zig");
const day5 = @import("day5/main.zig");
const day6 = @import("day6/main.zig");

const Log = zlog.Logger(.info, .{});
const ColorHandler = zlog.ColorHandler.Handler(.{ .timestamp = .rfc3339 });

pub fn main() !void {
    var handler = ColorHandler.init(zlog.stderr);
    var logger = try Log.init(.{ .handler = handler.handler(), .allocator = std.heap.page_allocator });
    defer logger.deinit();

    logger.info("advent of code 2015", .{});

    const allocator = std.heap.page_allocator;
    var args = try std.process.ArgIterator.initWithAllocator(allocator);
    defer _ = args.deinit();

    _ = args.next();

    if (args.next()) |day| {
        if (shouldRun(day, "1")) {
            try day1.part1();
            try day1.part2();
        }
        if (shouldRun(day, "2")) {
            try day2.part1();
            try day2.part2();
        }
        if (shouldRun(day, "3")) {
            try day3.part1();
            try day3.part2();
        }
        if (shouldRun(day, "4")) {
            try day4.part1();
            try day4.part2();
        }
        if (shouldRun(day, "5")) {
            try day5.part1();
            try day5.part2();
        }
        if (shouldRun(day, "6")) {
            try day6.part1();
            try day6.part2();
        }
    }
}

fn shouldRun(input: []const u8, day: []const u8) bool {
    if (std.mem.eql(u8, input, "all")) {
        return true;
    }
    if (std.mem.eql(u8, input, day)) {
        return true;
    }
    return false;
}
