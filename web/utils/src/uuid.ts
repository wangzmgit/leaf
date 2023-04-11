export const createUuid = () =>{
    const tmp_url = URL.createObjectURL(new Blob());
    const uuid = tmp_url.split('/').pop() as string;
    URL.revokeObjectURL(tmp_url);
    return uuid;
}