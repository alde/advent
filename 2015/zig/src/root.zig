pub const day1 = @import("day1/main.zig");
pub const day2 = @import("day2/main.zig");

test {
    @import("std").testing.refAllDecls(@This());
}
