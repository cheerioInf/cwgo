import { Col, Row } from "@douyinfe/semi-ui";
import styles from "./index.module.scss";

export default function AllInfoBox({ type }: { type: "idl" | "repo" }) {
	let data;
	if (type === "idl") {
		data = [
			{
				title: "recorded_psm_count",
				value: "1000"
			},
			{
				title: "normal_psm_count",
				value: "1000"
			},
			{
				title: "psm_list_last_full_sync_time",
				value: "2021-08-01 00:00:00"
			},
			{
				title: "idl_file_path_last_full_sync_time",
				value: "2021-08-01 00:00:00"
			},
			{
				title: "idl_info_last_full_sync_time",
				value: "2021-08-01 00:00:00"
			}
		];
	} else {
		data = [
			{
				title: "recorded_repository_count",
				value: "1000"
			},
			{
				title: "normal_repository_count",
				value: "1000"
			},
			{
				title: "repository_last_full_sync_time",
				value: "2021-08-01 00:00:00"
			},
			{
				title: "whether_to_stop_updating",
				value: "false"
			}
		];
	}

	return (
		<Row type="flex" justify="space-between" className={styles["all-info-box"]}>
			{data.map((item) => {
				return (
					<Col span={4}>
						<div className={styles["col-content"]}>
							<div
								style={{
									color: "var(--semi-color-text-1)",
									fontSize: "1rem"
								}}
							>
								{item.title}
							</div>
							<div>
								<strong
									style={{
										color: "var(--semi-color-primary)",
										fontSize: "1.5rem"
									}}
								>
									{item.value}
								</strong>
							</div>
						</div>
					</Col>
				);
			})}
		</Row>
	);
}