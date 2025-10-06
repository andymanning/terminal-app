-- Linux Boot Screen Simulation in Lua
-- This script simulates a typical Linux boot screen with green text on black background

-- Function to clear the screen
local function clear()
    io.write("\27[2J\27[H")
end

-- Function to set text color to green
local function green()
    io.write("\27[32m")
end

-- Function to reset text color
local function reset()
    io.write("\27[0m")
end

-- Function to print text with delay
local function print_with_delay(text, delay)
    io.write(text)
    io.flush()
    os.execute("sleep " .. delay)
end

-- Function to simulate loading progress
local function show_progress(width, delay_per_step)
    io.write("[")
    for i = 1, width do
        io.write("#")
        io.flush()
        os.execute("sleep " .. delay_per_step)
    end
    io.write("] Done\n")
end

-- Main boot sequence
clear()
green()

print_with_delay("Starting system boot...\n\n", 1)

print_with_delay("BIOS version 2.0.3 initialized\n", 0.5)
print_with_delay("CPU: Intel(R) Core(TM) i7-10700K @ 3.80GHz\n", 0.3)
print_with_delay("Memory: 16384MB (16GB)\n", 0.3)
print_with_delay("Initializing hardware...", 0.5)
show_progress(20, 0.05)

print_with_delay("Loading kernel...\n", 0.5)
print_with_delay("Linux version 5.15.0-generic (gcc version 11.2.0)\n", 0.3)
print_with_delay("Command line: BOOT_IMAGE=/boot/vmlinuz-5.15.0-generic root=UUID=1234-5678 ro quiet splash\n", 0.3)
print_with_delay("Loading initial ramdisk...", 0.5)
show_progress(30, 0.03)

print_with_delay("Mounting root filesystem...\n", 0.5)
print_with_delay("Checking filesystems...", 0.5)
show_progress(15, 0.04)

print_with_delay("Starting system services...\n", 0.5)
print_with_delay("[ OK ] Started System Logging Service.\n", 0.2)
print_with_delay("[ OK ] Started D-Bus System Message Bus.\n", 0.2)
print_with_delay("[ OK ] Started Network Manager.\n", 0.2)
print_with_delay("[ OK ] Started Accounts Service.\n", 0.2)
print_with_delay("[ OK ] Started CUPS Scheduler.\n", 0.2)
print_with_delay("[ OK ] Started Bluetooth service.\n", 0.2)
print_with_delay("[ OK ] Started Login Service.\n", 0.2)
print_with_delay("[ OK ] Started Disk Manager.\n", 0.2)
print_with_delay("[ OK ] Started System Security Services Daemon.\n", 0.2)
print_with_delay("[ OK ] Started User Manager for UID 1000.\n", 0.2)

print_with_delay("\nStarting desktop environment...\n", 0.5)
print_with_delay("Loading X server...", 0.5)
show_progress(25, 0.04)

print_with_delay("\nWelcome to Linux!\n", 1)
print_with_delay("Login: ", 0.5)

reset()
print_with_delay("\n\n(Boot screen simulation complete)\n", 0)