pub const day1 = @import("day1/main.zig");

test {
    @import("std").testing.refAllDecls(@This());
}
