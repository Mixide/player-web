<template>
  <Content>
    <div class="row">
        <div class="col-3">
          <UserInfo :user="user" />
          <Upload/>
        </div>
        <div class="col-9">
          <MusicList :list="music_list"/>
        </div>
      </div>
  </Content>
</template>

<script>
  import Content from '../components/Content';
  import UserInfo from '../components/UserInfo';
  import MusicList from '../components/MusicList';
  import Upload from '../components/Upload';
  import {reactive} from 'vue';
  export default {
    name: 'MusicSpace',
    components:{
      Content,
      UserInfo,
      MusicList,
      Upload,
    },
    setup() {
      const user = reactive({
        id: 1,
        username: "test",
        music_nums: 21,
      });
      const music_list = reactive({
        count:0,
        musics:[
          {
            id:1,
            name:"my_music1",
            singer: "Zhou",
            music_src:"file",
          },
          {
            id:2,
            name:"my_music2",
            singer: "Li",
            music_src:"file",
          }
        ]
      })
      return {user,music_list};
    }
  }
</script>

<style scoped>
  
</style>