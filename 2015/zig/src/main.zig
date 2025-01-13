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
        if (std.mem.eql(u8, day, "all") or std.mem.eql(u8, day, "1")) {
            try day1.part1();
            try day1.part2();
        }
        if (std.mem.eql(u8, day, "all") or std.mem.eql(u8, day, "2")) {
            try day2.part1();
            try day2.part2();
        }
        if (std.mem.eql(u8, day, "all") or std.mem.eql(u8, day, "3")) {
            try day3.part1();
            try day3.part2();
        }
        if (std.mem.eql(u8, day, "all") or std.mem.eql(u8, day, "4")) {
            try day4.part1();
            try day4.part2();
        }
    }
}
