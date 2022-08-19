const std = @import("std");
const io = std.io;
const fmt = std.fmt;

const BitCount = struct {
    zeros: u16,
    ones: u16,
};

const rll = 12;

pub fn main() !void {
    const stdin = io.getStdIn().reader();

    var buf: [128]u8 = undefined;
    var sum: [rll]BitCount = undefined;
    for (sum) |_,i| {
        sum[i] = BitCount{
            .zeros = 0,
            .ones = 0,
        };
    }

    // Read from stdin, line by line
    // TODO: sense rll to work with example or source input
    while (try stdin.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        for (line) |char, i| {
            if (char == '1') {
                sum[i].ones += 1;
            } else if (char == '0') {
                sum[i].zeros += 1;
            } else {
                std.debug.print("PARSING ERROR: 0x{d}", .{char});
            }
        }
    }
    var gamma_binary: [rll]u8 = undefined;
    var epsilon_binary: [rll]u8 = undefined;
    for (sum) |_,i| {
        if (sum[i].zeros > sum[i].ones){
            gamma_binary[i] = '0';
            epsilon_binary[i] = '1';
        } else {
            gamma_binary[i] = '1';
            epsilon_binary[i] = '0';
        }
    }
    const gamma = fmt.parseUnsigned(u32, gamma_binary[0..gamma_binary.len], 2) catch |err| {
        std.debug.print("PARSE ERR: {} {s}\n", .{ err, gamma_binary });
        return err;
    };
    const epsilon = fmt.parseUnsigned(u32, epsilon_binary[0..epsilon_binary.len], 2) catch |err| {
        std.debug.print("PARSE ERR: {} {s}\n", .{ err, gamma_binary });
        return err;
    };
    std.debug.print("gamma:   0b{s} {d}\n", .{gamma_binary, gamma});
    std.debug.print("epsilon: 0b{s} {d}\n", .{epsilon_binary, epsilon});
    std.debug.print("power:   {d}\n", .{epsilon * gamma});
}
