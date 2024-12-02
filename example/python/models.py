from py_go_bridge.python.bridge.types import PyGoBaseInput

class ExportExcelInput(PyGoBaseInput):
    """导出Excel的输入参数"""
    channel_code: str
    host: str = "localhost"
    port: int = 3306
    user: str = "root"
    password: str = "password"
    database: str = "test"
    worker_count: int = 4 


class EchoInput(PyGoBaseInput):
    input: str