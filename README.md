# Financial Management System

## 创建编译环境
```bash
cd deployment/compiler
docker build -t fms-compiler .
```

## 编译项目
```bash
cd /deployment/compiler
./build.sh
```

## 假数据生成
Usage
```bash
./faker --help
```

## 运行项目
```bash
cd deployment/compose
docker compose up
```