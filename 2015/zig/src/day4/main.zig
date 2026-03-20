const std = @import("std");
const shared = @import("../shared.zig");

const hash = std.crypto.hash;
const alloc = std.heap.page_allocator;

pub fn part1() !void {
    var logger = try shared.initLogger();
    defer logger.deinit();
    logger.info("starting", .{ .day = 4, .part = 1 });
    const result = try firstDigit("ckczppom", 5);
    logger.info("completed", .{ .day = 4, .part = 1, .result = result });
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
    var logger = try shared.initLogger();
    defer logger.deinit();
    logger.info("starting", .{ .day = 4, .part = 2 });
    const result = try firstDigit("ckczppom", 6);
    logger.info("completed", .{ .day = 4, .part = 2, .result = result });
}

test "day 4" {
    const testing = std.testing;

    {
        try testing.expectEqual(609043, firstDigit("abcdef", 5));
        try testing.expectEqual(1048970, firstDigit("pqrstuv", 5));
    }

    {}
}
