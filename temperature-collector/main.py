from http.server import BaseHTTPRequestHandler, HTTPServer
import time
from datetime import datetime

TEMP_FILE_PATH = "/sys/class/thermal/thermal_zone0/temp"

data = 0.0
updated_at = datetime.now()

class Handler(BaseHTTPRequestHandler):
  def do_GET(self):
    global data

    self.send_response(200)
    self.end_headers()
    self.wfile.write('# HELP device_temperature Temperature measured by lm sensors\n'.encode())
    self.wfile.write('# TYPE device_temperature gauge\n'.encode())
    self.wfile.write(f'device_temperature {data}\n'.encode())
    
import threading
def regular_update():
  global data
  while True:
    f = open(TEMP_FILE_PATH, 'r')
    data = int(f.read())
    f.close()
    updated_at = datetime.now()
    time.sleep(10)


def start_server():
  server = HTTPServer(('0.0.0.0', 8000), Handler)
  try:
    print("start server")
    server.serve_forever()
  except KeyboardInterrupt:
    pass

  server.server_close()
  print("Server stopped.")

if __name__ == "__main__":
  thread = threading.Thread(target=regular_update, daemon=True)
  thread.start()
  start_server()

# sample response
'''
# HELP device_temperature Temperature measured by lm sensors
# TYPE device_temperature gauge
device_temperature 5.0
'''