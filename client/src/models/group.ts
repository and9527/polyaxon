export class GroupModel {
	public uuid: string;
	public unique_name: string;
	public sequence: number;
	public num_experiments: number;
	public project_name: string;
	public user: string;
	public concurrency: number;
	public content: number;
	public num_pending_experiments: number;
	public num_running_experiments: number;
	public deleted?: boolean;
	public createdAt: Date;
	public updatedAt: Date;
	public started_at: Date;
	public finished_at: Date;
}

export class GroupStateSchema {
	byUuids: {[uuid: string]: GroupModel};
	uuids: string[];
}

export const GroupsEmptyState = {byUuids: {}, uuids: []};
