pub const shared = @import("shared.zig");
pub const day1 = @import("day1/main.zig");
pub const day2 = @import("day2/main.zig");
pub const day3 = @import("day3/main.zig");
pub const day4 = @import("day4/main.zig");
pub const day5 = @import("day5/main.zig");

test {
    @import("std").testing.refAllDecls(@This());
}
