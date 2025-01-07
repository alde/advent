const std = @import("std");
const fs = std.fs;

/// Reads the entire content of a file as a single long line.
/// Returns the content as a `[]const u8`.
pub fn readInput(file_path: []const u8) ![]const u8 {
    const allocator = std.heap.page_allocator;

    const file = try std.fs.cwd().openFile(file_path, .{});
    defer file.close();

    const file_size = try file.getEndPos();
    const buffer = try allocator.alloc(u8, file_size);

    const read_count = try file.readAll(buffer);
    if (read_count != file_size) {
        return error.UnexpectedEOF;
    }

    return buffer;
}
