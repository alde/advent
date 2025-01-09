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

const ReaderType = std.fs.File.Reader;
const BufReaderType = std.io.BufferedReader(4096, ReaderType);
const BufReaderReaderType = BufReaderType.Reader;

pub const ReadByLineIterator = struct {
    file: std.fs.File,
    reader: ReaderType,
    buf_reader: BufReaderType,
    stream: ?BufReaderReaderType,
    buf: [4096]u8,

    pub fn next(self: *@This()) !?[]u8 {
        if (self.stream == null) {
            self.stream = self.buf_reader.reader();
        }
        if (self.stream) |stream| {
            return stream.readUntilDelimiterOrEof(&self.buf, '\n');
        }
        unreachable;
    }

    pub fn deinit(self: *@This()) void {
        self.file.close();
    }
};

// Iterate over the lines in the file using a buffered reader.
// Caller is responsible for calling deinit() on returned iterator when done.
pub fn iterLines(filename: []const u8) !ReadByLineIterator {
    var file = try std.fs.cwd().openFile(filename, .{});
    const reader = file.reader();
    const buf_reader = std.io.bufferedReader(reader);

    return ReadByLineIterator{
        .file = file,
        .reader = reader,
        .buf_reader = buf_reader,
        .stream = null,
        .buf = undefined,
    };
}
