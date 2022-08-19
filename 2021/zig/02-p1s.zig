const std = @import("std");
const io = std.io;
const fmt = std.fmt;

const Vector = enum { forward, up, down };

// Submarine Commands
const Command = struct {
    vector: Vector,
    thrust: u8,
};

const Coordinate = struct {
    depth: u16,
    forward: u16,
};

pub fn main() !void {
    const stdin = io.getStdIn().reader();
    const allocator = std.heap.page_allocator;

    // Read from stdin, line by line
    // split on " ", parse, and
    // insert into ArrayList of Commands
    var buf: [128]u8 = undefined;
    var course = std.ArrayList(Command).init(allocator);
    defer course.deinit();
    while (try stdin.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var line_parts = std.mem.split(u8, line, " ");
        var c = Command{
            .vector = undefined,
            .thrust = undefined,
        };
        const raw_vector = line_parts.first();
        if (std.mem.eql(u8, raw_vector, "forward")) {
            c.vector = Vector.forward;
        } else if (std.mem.eql(u8, raw_vector, "up")) {
            c.vector = Vector.up;
        } else if (std.mem.eql(u8, raw_vector, "down")) {
            c.vector = Vector.down;
        }
        var raw_thrust = line_parts.next().?;
        c.thrust = fmt.parseUnsigned(u8, raw_thrust, 10) catch |err| {
            std.debug.print("PARSE ERR: {} {s}\n", .{ err, raw_thrust });
            continue;
        };

        try course.append(c);
        //std.debug.print("{}\n", .{c});
    }
    var p = Coordinate{
        .depth = 0,
        .forward = 0,
    };
    const cmds = course.items;
    for (cmds) |c| {
        p = switch (c.vector) {
            .forward => Coordinate{ .depth = p.depth, .forward = p.forward + c.thrust },
            .up => Coordinate{ .depth = p.depth - c.thrust, .forward = p.forward },
            .down => Coordinate{ .depth = p.depth + c.thrust, .forward = p.forward },
        };
    }
    const ans: u32 = @as(u32, p.depth) * p.forward;
    std.debug.print("{}, {d}\n", .{ p, ans });
}
