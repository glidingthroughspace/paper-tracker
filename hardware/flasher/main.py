import subprocess
import os
import traceback
import json
import PySimpleGUI as sg

sg.theme('DarkAmber')

TITLE = 'Paper-Tracker Firmware Flasher'
KEY_FW_DIR = 'fw-dir'
KEY_SSID = 'wifi-ssid'
KEY_USERNAME = 'wifi-user'
KEY_PASSWORD = 'wifi-pass'
KEY_IP = 'server-ip'
KEY_PORT = 'device-port'
KEY_FLASH_OUTPUT = 'flash-output'
KEY_CLOSE = 'flash-close'

def getPIODevices():
    raw_results = subprocess.check_output(['pio', 'device', 'list', '--json-output'], text=True)
    results = json.loads(raw_results)
    print(results)
    return tuple(result['port']+' ('+result['description']+')' for result in results)

def generateCredentials(values):
    example_content = None
    with open(values[KEY_FW_DIR]+'/include/credentials.example.hpp', 'r') as example_file:
        example_content = example_file.read()

    content = example_content.replace('$$WIFI_SSID$$', values[KEY_SSID])
    content = content.replace('$$WIFI_USERNAME$$', values[KEY_USERNAME])
    content = content.replace('$$WIFI_PASSWORD$$', values[KEY_PASSWORD])
    ip = values[KEY_IP].split('.')
    content = content.replace('$$IP_1$$', ip[0])
    content = content.replace('$$IP_2$$', ip[1])
    content = content.replace('$$IP_3$$', ip[2])
    content = content.replace('$$IP_4$$', ip[3])

    with open(values[KEY_FW_DIR]+'/include/credentials.hpp', 'w') as file:
        file.write(content)

def flash(values, window, output):
    port = values[KEY_PORT].split(' ')[0]
    cmd = ['pio', 'run', '-e', 'tinypico', '-t', 'upload', '--upload-port', port]
    process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True ,cwd=values[KEY_FW_DIR])
    while True:
        if process.poll() is not None:
            break
        line = process.stdout.readline()
        if line:
            output.Update(value='\t', append=True)
            output.Update(value=line, append=True)
            window.Refresh()
    rc = process.poll()
    if rc != 0:
        raise subprocess.CalledProcessError(rc, cmd)

pio_devices = getPIODevices()
input_layout = [  [sg.Text(TITLE, font='any 20')],
            [sg.Text('Firmware Directory:', size=(15, 1)), sg.Input(key=KEY_FW_DIR), sg.FolderBrowse()],
            [sg.Text('WiFi SSID:', size=(15, 1)), sg.InputText(key=KEY_SSID)],
            [sg.Text('WiFi Username:', size=(15, 1)), sg.InputText(key=KEY_USERNAME)],
            [sg.Text('WiFi Password:', size=(15, 1)), sg.InputText(key=KEY_PASSWORD)],
            [sg.Text('Server IP:', size=(15, 1)), sg.InputText(key=KEY_IP)],
            [sg.Text('Port', size=(15, 1)), sg.Combo(values=pio_devices, default_value=pio_devices[0], size=(40, 1), key=KEY_PORT)],
            [sg.Button('Flash', size=(15, 1)), sg.Button('Cancel')] ]

flash_layout = [  [sg.Text(TITLE, font='any 20')],
            [sg.Text('Flashing...')],
            [sg.MLine(key=KEY_FLASH_OUTPUT, size=(100, 15), autoscroll=True)],
            [sg.Button('Close', key=KEY_CLOSE, disabled=True)] ]

def main():
    input_window = sg.Window(TITLE, input_layout)

    event, values = input_window.read()
    input_window.close()

    if event in (None, 'Cancel'):
        return

    flash_window = sg.Window(TITLE, flash_layout, finalize=True)
    output = flash_window.FindElement(KEY_FLASH_OUTPUT)
    close = flash_window.FindElement(KEY_CLOSE)

    try:
        output.Update(value='Generate Credentials File...\n', append=True)
        flash_window.Refresh()
        generateCredentials(values)
        output.Update(value='...Done\n', append=True)
        flash_window.Refresh()

        output.Update(value='Flash firmware...\n\n', append=True)
        flash_window.Refresh()
        flash(values, flash_window, output)
        output.Update(value='\n\n...Done\n\nFlashing finished!', append=True)
        flash_window.Refresh()
    except Exception as e:
        traceback.print_exc()
        text = '..Failed: ' + str(e) + '\n'
        output.Update(value=text, append=True)
    finally:
        close.Update(disabled=False)

    flash_window.read()

if __name__ == "__main__":
    main()
