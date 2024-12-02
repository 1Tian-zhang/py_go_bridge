from pydantic import BaseModel

class PyGoBaseInput(BaseModel):
    """所有Go函数输入参数的基类"""
    def to_c_args(self) -> list:
        """将Pydantic模型转换为C参数列表"""
        args = []
        for field_name in self.model_fields:
            value = getattr(self, field_name)
            if isinstance(value, str):
                args.append(value.encode('utf-8'))
            else:
                args.append(value)
        return args

class NewFunctionInput(PyGoBaseInput):
    """新函数的输入参数"""
    param1: str
    param2: int

class ExportExcelInput(PyGoBaseInput):
    """导出Excel的输入参数"""
    channel_code: str
    host: str = "localhost"
    port: int = 3306
    user: str = "root"
    password: str = "password"
    database: str = "test"
    worker_count: int = 4 