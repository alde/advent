const std = @import("std");
const zlog = @import("zlog");

pub const Log = zlog.Logger(.info, .{});

const ColorHandler = zlog.ColorHandler.Handler(.{ .timestamp = .rfc3339 });
var handler = ColorHandler.init(zlog.stderr);

pub fn initLogger() !Log {
    const allocator = std.heap.page_allocator;
    return Log.init(.{ .handler = handler.handler(), .allocator = allocator });
}

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

pub const ReadByLineIterator = struct {
    file: std.fs.File,
    file_reader: std.fs.File.Reader,

    pub fn next(self: *@This()) !?[]u8 {
        return self.file_reader.interface.takeDelimiter('\n');
    }

    pub fn deinit(self: *@This()) void {
        self.file.close();
    }
};

/// Iterate over the lines in the file using a buffered reader.
/// Caller is responsible for calling deinit() on returned iterator when done.
pub fn iterLines(filename: []const u8) !ReadByLineIterator {
    const allocator = std.heap.page_allocator;
    const file = try std.fs.cwd().openFile(filename, .{});
    const buf = try allocator.alloc(u8, 4096);
    const file_reader = file.reader(buf);

    return ReadByLineIterator{
        .file = file,
        .file_reader = file_reader,
    };
}

pub fn stringContains(haystack: []const u8, needle: []const u8) bool {
    var i: usize = 0;
    if (haystack.len < needle.len) {
        return false;
    }

    if (haystack.len == needle.len) {
        return std.mem.eql(u8, haystack, needle);
    }
    const n = needle.len;
    while (i < haystack.len) : (i += 1) {
        if (i + n > haystack.len) {
            break;
        }
        if (std.mem.eql(u8, haystack[i .. i + n], needle)) {
            return true;
        }
    }
    return false;
}
test "stringContains" {
    const testing = std.testing;

    try testing.expectEqual(true, stringContains("abba", "ab"));
    try testing.expectEqual(true, stringContains("abba", "bb"));
    try testing.expectEqual(true, stringContains("abba", "ba"));
    try testing.expectEqual(false, stringContains("abba", "bab"));
}
