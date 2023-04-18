# SS Test Endpoint with Postman

## scenario umum
### Register
Method : POST
endpoint : http://localhost:PORT/users/register
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681831335/FGA%20GO%2010/challange-09/image_2023-04-18_222218622_jvwurf.png)

### login
Method : POST
endpoint : http://localhost:PORT/users/login
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681831426/FGA%20GO%2010/challange-09/image_2023-04-18_222351643_ojv5nt.png)

### protected route (auth required)
akan memunculkan respon unauthenthicated kalau belum login (memasukan barier token pada header auth)
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681831542/FGA%20GO%2010/challange-09/image_2023-04-18_222547635_lqcrra.png)

## scenario user biasa
### create produk
Method : POST
endPoint : http://localhost:PORT/products
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681831653/FGA%20GO%2010/challange-09/image_2023-04-18_222739141_jnli5k.png)

### Read produk
method : GET
endpoint : http://localhost:PORT/products/:productId
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681831744/FGA%20GO%2010/challange-09/image_2023-04-18_222909631_qlmeng.png)

### tapi ketika liat produk user lain
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832351/FGA%20GO%2010/challange-09/image_2023-04-18_223914898_dtrgwz.png)

### lalu akses route yang hanya bisa diakses admin
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832451/FGA%20GO%2010/challange-09/image_2023-04-18_224056215_hhe6g6.png)

## scenario admin
### create produk
method : post
endpoint : http://localhost:PORT/products
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832567/FGA%20GO%2010/challange-09/image_2023-04-18_224252701_zob1ak.png)

### update produk
method : PUT
endpoint : http://localhost:PORT/products/:productId
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832644/FGA%20GO%2010/challange-09/image_2023-04-18_224410375_jfu1td.png)

### read produk by id
method : GET
endpoint : http://localhost:PORT/products/:productId
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832710/FGA%20GO%2010/challange-09/image_2023-04-18_224515520_cdct66.png)

### read produk semua
method : GET
endpoint : http://localhost:PORT/products
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832825/FGA%20GO%2010/challange-09/image_2023-04-18_224711118_xcyqhz.png)

### delete product
method : delete
endpoint : http://localhost:PORT/products/:productId
![image](https://res.cloudinary.com/drakr4vtu/image/upload/v1681832874/FGA%20GO%2010/challange-09/image_2023-04-18_224800412_oaodv5.png)

