"""Idler - Simple Python program to prevent AFK timeouts during FFXIV Endwalker launch."""
#!/usr/bin/python3

from datetime import datetime
import time
import keyboard

def main():
    """Main"""
    last_keypress_time = time.time()

    print("Idler - presses escape 4x after any 10 minute interval of no user input.")
    print("Idler started...")

    # Setup keyboard event hook
    keyboard.on_press(key_event)

    # Loop to calc time since last keypress - every 1 second
    while True:
        if time.time() - last_keypress_time >= 600:
            curr_datetime = datetime.now().strftime("%H:%M:%S")
            print(f"Idler activated at {curr_datetime}, pressing keys...")
            for _ in range(4):
                keyboard.press_and_release('escape')
                time.sleep(1)

            # Set macro keypress to be most recent
            last_keypress_time = time.time()

        time.sleep(1)

def key_event(keypress: keyboard.KeyboardEvent):
    """Keyboard event callback function"""
    global last_keypress_time

    # Set last keypress time to current time
    # Ignore escape keypresses, as those are
    # reserved for macro keypresses by the script
    if keypress.scan_code != 53:
        last_keypress_time = time.time()
        print("Key pressed: " + str(keypress.name))

main()
