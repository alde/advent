const std = @import("std");
const day1 = @import("day1/main.zig");
const day2 = @import("day2/main.zig");
const day3 = @import("day3/main.zig");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    try day1.part1();
    try day1.part2();
    try stdout.print("*-*-*-*-*-*-*-*-*\n", .{});
    try day2.part1();
    try day2.part2();
    try stdout.print("*-*-*-*-*-*-*-*-*\n", .{});
    try day3.part1();
    try day3.part2();
}
