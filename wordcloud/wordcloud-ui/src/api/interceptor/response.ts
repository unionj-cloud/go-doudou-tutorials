export const responseUse = (response:any) => {
  if(response.status==200){
    return response.data
  }
  else return response
}
