const std = @import("std");
const shared = @import("../shared.zig");

const hash = std.crypto.hash;
const alloc = std.heap.page_allocator;

pub fn part1() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 4 Part 1\n", .{});
    const result = try firstDigit("ckczppom", 5);
    try stdout.print("Result: {}\n", .{result});
}

fn firstDigit(input: []const u8, numZeros: u8) !usize {
    var padding: usize = 1;
    while (true) : (padding += 1) {
        const candidate = try std.fmt.allocPrint(alloc, "{s}{d}", .{ input, padding });
        defer alloc.free(candidate);

        var out: [hash.Md5.digest_length]u8 = undefined;
        hash.Md5.hash(candidate, &out, .{});

        if (out[0] == 0 and out[1] == 0) {
            if ((numZeros == 5 and out[2] <= 15) or (numZeros == 6 and out[2] == 0)) {
                return padding;
            }
        }
    }
}

pub fn part2() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Day 4 Part 2\n", .{});
    const result = try firstDigit("ckczppom", 6);

    try stdout.print("Result: {}\n", .{result});
}

test "day 4" {
    const testing = std.testing;

    {
        try testing.expectEqual(609043, firstDigit("abcdef"));
        try testing.expectEqual(1048970, firstDigit("pqrstuv"));
    }

    {}
}
