import time

if __name__ == '__main__':
    time_stamp = int(time.time())
    if time_stamp % 2 == 0:
        raise Exception(print("ops!"))
