<template>
  <div class="card list">
    <div class="card-body">
      <div v-for="(music, index) in list.musics" :key="index" @click="handleCardClick(index)">
        <div class="card music">
          <div class="card-body">
            <div class="row">
              <div class="col-10 music_name">{{ music.name }}</div>
              <div class="col-2 artist">歌手：{{ music.artist }}</div>
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
      },
      ap:{
        type:Object,
        required:true
      }
    },
    watch: {
    // 监听 musicList prop 的变化
      ap: {}
    },
    methods: {
      handleCardClick(index) {
        this.ap.value.pause()
        this.ap.value.list.switch(index);
        this.ap.value.play();
      }
    }
  }
</script>

<style scoped>
.singer {
    font-size: 14px;
    color: gray;
}  
.list {
  height: 620px;
}
.music {
  margin-top: 10px;
  cursor: pointer;
}
.music:hover {
  box-shadow: 2px 2px 10px lightgray;
  transition: 300ms;
}
</style>