import {
	Button,
	ConfigProvider,
	Modal,
	Space,
	Table,
	Toast
} from "@douyinfe/semi-ui";
import { getTempItem } from "../api";
import { useEffect, useState } from "react";
import en_GB from "@douyinfe/semi-ui/lib/es/locale/source/en_GB";
import CodeMirror from "@uiw/react-codemirror";
import { okaidia } from "@uiw/codemirror-theme-okaidia";
import { Data } from "@douyinfe/semi-ui/lib/es/table";
// import { yaml } from "@codemirror/legacy-modes/mode/yaml";

const pageSize = 5;

export default function InnerTable() {
	const data = getTempItem();
	const [dataSource, setData] = useState<unknown>([]);
	const [loading, setLoading] = useState(false);
	const [currentPage, setPage] = useState(1);
	const [modal, contextHolder] = Modal.useModal();
	const config = {
		width: "50%",
		title: "查看 / 更新配置文件",
		content: (
			<div>
				<CodeMirror
					value={`
path: main.go
update_behavior:
type: cover
body: |-
/*
* Copyright (c) 2023 lanshan team. All rights reserved.
*/

package main

import (
"fmt"
apollo "wecqupt/app/common/config"
"wecqupt/app/common/config/model"
"{{.ImportPath}}/{{ToLower .ServiceName}}"
"wecqupt/app/common/logger"
"{{.Module}}/internal/config"
"{{.Module}}/internal/svc"
"wecqupt/utils"

"go.uber.org/zap"
)

var c config.Config

func main() {
// 初始化配置管理器
err := apollo.InitClient(config.ServiceName)
if err != nil {
    panic(fmt.Sprintf("initialize Apollo Client failed, err: %v", err))
}

namespace, serviceSingleName := utils.GetServiceDetails(config.ServiceName)

err = apollo.Common().UnmarshalServiceConfig(namespace, serviceSingleName, &c)

// 初始化服务
c.SetUp()

// 初始化服务上下文
svcCtx := svc.NewServiceContext(c)

// 启动服务
svr := {{ToLower .ServiceName}}.NewServer(
    &{{.ServiceName}}Impl{
        svcCtx: svcCtx,
    },
    model.KitexServerOptions...,
)

err = svr.Run()
if err != nil {
    logger.Logger.Error("kitex server run failed.", zap.Error(err))
}
}`}
					height="500px"
					theme={okaidia}
					// extensions={[yaml()]}
					// onChange={onChange}
				/>
			</div>
		),
		cancelText: "取消",
		okText: "确定",
		icon: null,
		onOk: () => {
			// 返回一个延时的 Promise
			return new Promise((resolve, reject) => {
				setTimeout(
					Math.random() > 0.5
						? () => {
								Toast.success("添加成功！");
								resolve(true);
						  }
						: () => {
								Toast.error("Oops errors!");
								reject(false);
						  },
					1000
				);
			}).catch(() => console.log("Oops errors!"));
		}
	};

	const columns = [
		{
			title: "模版元素名称",
			dataIndex: "name",
			width: 250,
			render: (value: string) => {
				return <div>{value}</div>;
			}
		},
		{
			title: "创建时间",
			dataIndex: "create_time",
			width: 350,
			render: (value: string) => {
				return <div>{value}</div>;
			}
		},
		{
			title: "更新时间",
			dataIndex: "update_time",
			width: 350,
			render: (value: string) => {
				return <div>{value}</div>;
			}
		},
		{
			title: "操作",
			render: (value: string) => {
				console.log("value", value);
				return (
					<Space>
						<Button
							onClick={() => {
								console.log("modal", modal);
								modal.confirm(config);
							}}
						>
							查看 / 更新配置
						</Button>
					</Space>
				);
			}
		}
	];

	const fetchData = async (currentPage = 1) => {
		setLoading(true);
		setPage(currentPage);
		const curDataSource = await new Promise((res) => {
			setTimeout(() => {
				const data = getTempItem();
				const dataSource = data.slice(
					(currentPage - 1) * pageSize,
					currentPage * pageSize
				);
				res(dataSource);
			}, 300);
		});
		setLoading(false);
		setData(curDataSource);
	};

	const handlePageChange = (page: number) => {
		fetchData(page);
	};

	useEffect(() => {
		fetchData();
	}, []);

	return (
		<ConfigProvider locale={en_GB}>
			<Table
				columns={columns}
				dataSource={dataSource as Data[]}
				pagination={{
					currentPage,
					pageSize: 5,
					total: data.length,
					onPageChange: handlePageChange
				}}
				loading={loading}
			/>
			{contextHolder}
		</ConfigProvider>
	);
}