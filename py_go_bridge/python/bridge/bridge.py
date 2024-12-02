import ctypes
import json
import os
from typing import Any
from py_go_bridge.python.bridge.models import PyGoBaseInput

class Bridge:
    def __init__(self, lib_path: str):
        """初始化桥接实例"""
        if not os.path.exists(lib_path):
            raise ValueError(f"Library not found: {lib_path}")
            
        self.lib = ctypes.CDLL(lib_path)
        self._init_free_string()
        
    def _init_free_string(self):
        """初始化释放字符串的函数"""
        self.lib.FreeString.argtypes = [ctypes.POINTER(ctypes.c_char)]
        self.lib.FreeString.restype = None
        
    def call_go_function(self, func_name: str, input_data: PyGoBaseInput) -> dict:
        """调用Go函数"""
        # 获取函数
        func = getattr(self.lib, func_name)
        func.restype = ctypes.POINTER(ctypes.c_char)
        
        # 准备参数
        args = input_data.to_c_args()
        
        # 调用函数
        result_ptr = func(*args)
        
        try:
            if not result_ptr:
                raise Exception("Null result from Go function")
                
            # 转换结果
            result_str = ctypes.cast(result_ptr, ctypes.c_char_p).value
            return json.loads(result_str)
            
        finally:
            if result_ptr:
                self.lib.FreeString(result_ptr)
                
    def __getattr__(self, name: str):
        """动态处理函数调用"""
        def wrapper(input_data: PyGoBaseInput) -> Any:
            # 转换函数名(snake_case to PascalCase)
            go_func_name = ''.join(word.capitalize() for word in name.split('_'))
            return self.call_go_function(go_func_name, input_data)
        return wrapper 