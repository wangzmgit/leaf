export interface PartitionType {
    id: number,
    content: string
    parent_id: number
}

export interface AddPartitionType {
    content: string
    parentId: number
}