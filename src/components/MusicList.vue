<template>
  <div class="card">
    <div class="card-body">
      <div v-for="music in list.musics" :key="music.id">
        <div class="card">
          <div class="card-body">
            <div class="row">
              <div class="col-10 music_name">{{ music.name }}</div>
              <div class="col-2 singer">歌手：{{ music.singer }}</div>
            </div>
          </div>
        </div>
      </div>  
    </div>
  </div>
</template>

<script>
  export default{
    name:"MusicList",
    props:{
      list:{
        type:Object,
        required:true
      }
    }
  }
</script>

<style scoped>
.singer {
    font-size: 14px;
    color: gray;
}  
</style>