<template>
  <div class="card">
    <div class="card-body">
      <div class="row">
        <div class="col-3">
          <img class="img-fluid" src="../../server/assets/logo.png">
        </div>
        <div class="col-9">
          <div class="username">{{ user.username }}</div>
          <div class="music-nums">音乐数:{{ user.music_nums }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default{
    name: "UserInfo",
    props:{
        user:{
            type:Object,
            required:true,
        }
    }
}
</script>

<style scoped>
img {
    border-radius: 50%;
}
.username {
    font-weight: bold;
}
.music-nums{
    font-size: 12px;
    color: gray;
}
button {
    padding: 2px 4px;
    font-size: 12px;
}
</style>