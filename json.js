
arr=[
    'C1105 KNNL-CO-MP-C2-199-ALA-2023-24 SLAO BAGALKOT.pdf',
    'G0366 KNN-CO-MP-C2-148-ALA-2023-24 SLAO BAGALKOT.pdf',
    'G0431 KNN-CC-GP-C2-133-ALA-2023-24 SLAO BAGALKOT.pdf',
    'G0485 KNN-CO-C2-MP-162-ALA-2023-24 SLAO BAGALKOT.pdf',
    'G0571 KNN-CO-C2-GP-161-ALA-2023-24 SLAO BAGALKOT.pdf',
    'G0572 KNN-CO-C2-GP-173-ALA-2023-24 SLAO BAGALKOT.pdf',
    'G0573 KNN-CO-C2-ALA-149-GP-2023-24 SLAO BAGALKOT.pdf',
    
]

function arrObj(arr){
    arr.map(e=>convert(e))
}
function convert(s){
obj = {File_Name:"",Code:"",Year:"",Office:2}
obj["File_Name"]=s
obj["Code"]=s.split(" ")[0]
obj["Year"]=arr[1].split(" ")[1].slice(-7)
console.log(obj,',')
res.push(obj)
}
res=[]
arrObj(arr)

console.log(res.length)