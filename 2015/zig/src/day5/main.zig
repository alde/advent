const std = @import("std");
const shared = @import("../shared.zig");

const stdout = std.io.getStdOut().writer();

pub fn part1() !void {
    try stdout.print("Day 5 Part 1\n", .{});
    var line_it = try shared.iterLines("src/day5/input.txt");
    defer line_it.deinit();

    var result: i32 = 0;
    while (try line_it.next()) |line| {
        if (isStringNice(line)) {
            result += 1;
        }
    }
    try stdout.print("Result: {}\n", .{result});
}

fn isStringNice(input: []const u8) bool {
    var vowelCount: i32 = 0;
    var hasRepeat: bool = false;
    for (input, 0..) |c, i| {
        if (c == 'a' or c == 'e' or c == 'i' or c == 'o' or c == 'u') {
            vowelCount += 1;
        }
        // doing this here to check if the last character is a vowel
        if (i > input.len - 2) {
            break;
        }
        if ((c == 'a' and input[i + 1] == 'b') or (c == 'c' and input[i + 1] == 'd') or (c == 'p' and input[i + 1] == 'q') or (c == 'x' and input[i + 1] == 'y')) {
            return false;
        }
        if (c == input[i + 1]) {
            hasRepeat = true;
        }
    }
    if (hasRepeat and vowelCount >= 3) {
        return true;
    }
    return false;
}

pub fn part2() !void {
    try stdout.print("Day 5 Part 2\n", .{});
    var line_it = try shared.iterLines("src/day5/input.txt");
    defer line_it.deinit();

    var result: i32 = 0;
    while (try line_it.next()) |line| {
        if (isNiceImproved(line)) {
            result += 1;
        }
    }
    try stdout.print("Result: {}\n", .{result});
}

fn isNiceImproved(input: []const u8) bool {
    var hasRepeatedPair = false;
    var i: usize = 1;
    while (i < input.len) : (i += 1) {
        const pair: []const u8 = &[2]u8{ input[i - 1], input[i] };
        if (shared.stringContains(input[i + 1 ..], pair)) {
            hasRepeatedPair = true;
        }
    }

    var hasRepeatWithPadding = false;
    for (0..input.len - 2) |idx| {
        if (input[idx] == input[idx + 2]) {
            hasRepeatWithPadding = true;
            break;
        }
    }
    return hasRepeatedPair and hasRepeatWithPadding;
}

test "day 5 part one" {
    try std.testing.expectEqual(true, isStringNice("ugknbfddgicrmopn"));
    try std.testing.expectEqual(true, isStringNice("aaa"));
    try std.testing.expectEqual(false, isStringNice("jchzalrnumimnmhp"));
    try std.testing.expectEqual(false, isStringNice("haegwjzuvuyypxyu"));
    try std.testing.expectEqual(false, isStringNice("dvszwmarrgswjxmb"));
}

test "day 5 part two" {
    try std.testing.expectEqual(true, isNiceImproved("qjhvhtzxzqqjkmpb"));
    try std.testing.expectEqual(true, isNiceImproved("xxyxx"));
    try std.testing.expectEqual(false, isNiceImproved("aaa"));
    try std.testing.expectEqual(false, isNiceImproved("aaabcb"));
    try std.testing.expectEqual(false, isNiceImproved("uurcxstgmygtbstg"));
    try std.testing.expectEqual(false, isNiceImproved("ieodomkazucvgmuy"));
}
