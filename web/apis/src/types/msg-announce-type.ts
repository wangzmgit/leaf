export interface AnnounceType {
    id: number,
    title: string,
    content: string,
    created_at: string,
    url: string,
    important: boolean
}


export interface AddAnnounceType {
    title: string,
    content: string,
    url: string,
    important: boolean
}