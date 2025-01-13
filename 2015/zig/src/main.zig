const std = @import("std");
const day1 = @import("day1/main.zig");
const day2 = @import("day2/main.zig");
const day3 = @import("day3/main.zig");
const day4 = @import("day4/main.zig");

pub fn main() !void {
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
