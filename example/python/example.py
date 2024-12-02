from py_go_bridge.python.bridge.bridge import Bridge
from models import ExportExcelInput, EchoInput


def main():
    bridge = Bridge("../go/bridge.so")
    
    result = bridge.export_excel(ExportExcelInput(
        channel_code="test",
        host="localhost",
        port=3306,
        user="root",
        password="hashchat",
        database="test",
        worker_count=4
    ))
    print(result)
    echo_input = EchoInput(input="hello")
    result = bridge.echo(echo_input)
    print(result)

if __name__ == "__main__":
    main() 