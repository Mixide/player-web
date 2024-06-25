<template>
  <div class="block text-center" style="margin-top: 30px; margin-bottom: 20px; margin-left: 100px; margin-right: 100px;">
    <!-- <span class="demonstration">Motion blur the switch (default)</span> -->
    <el-carousel :interval="4000" type="card" height="400px" width="800px">
      <el-carousel-item v-for="item in images" :key="item.id">
        <img :src="item.imageUrl" alt="Image" width="800" height="400">
        <!-- <img src="../assets/1.jpg" alt="Logo" width="800" height="400">
        <img src="../assets/2.jpg" alt="Logo" width="800" height="400"> -->
        <div class="image-overlay"></div>
        <p class="carousel-text">{{ item.caption }}</p>
      </el-carousel-item>
    </el-carousel>
  </div>
  <!-- <p class="text-center demonstration">Vertical effect</p> -->
  
</template>

<script>
export default {
  data(){
    return {
      images: [
        { id: 1, imageUrl: require('@/assets/1.jpg'), caption: "Hello" },
        { id: 2, imageUrl: require('@/assets/2.jpg'), caption: "Hello" },
        { id: 3, imageUrl: require('@/assets/3.jpg'), caption: "Hello" },
        { id: 4, imageUrl: require('@/assets/4.jpg'), caption: "Hello" },
      ],
    };
  },
}

</script>

<style scoped>
.demonstration {
  color: var(--el-text-color-secondary);
}
.image-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 120px;
  background: linear-gradient(to bottom, rgba(0,0,0,0) 0%, rgba(0,0,0,0.8) 100%);
  filter: blur(10px);
}
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 400px;
  width : 800px;
  margin: 0;
  text-align: center;
  background-color: #99a9bf;
}
.carousel-text {
  position: absolute;
  bottom: 20px;
  left: 10%;
  transform: translateX(-50%);
  color: white;
  font-weight: bold;
  font-size: 22px;
  text-align: left;
}

/* .el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
} */
</style>
