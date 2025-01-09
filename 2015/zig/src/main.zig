const std = @import("std");
const day1 = @import("day1/main.zig");
const day2 = @import("day2/main.zig");

pub fn main() !void {
    try day1.part1();
    try day1.part2();

    try day2.part1();
    try day2.part2();
}
