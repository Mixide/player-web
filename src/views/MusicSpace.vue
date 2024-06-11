<template>
  <AudioBar v-if="music_list.musics.length > 0" :audio="music_list.musics"/>
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
  import AudioBar from '../components/AudioBar';
  import {onMounted, reactive, ref} from 'vue';
  import axios from 'axios';
  export default {
    name: 'MusicSpace',
    components:{
      Content,
      UserInfo,
      MusicList,
      Upload,
      AudioBar,
    },
    setup() {
      const user = ref({
        id: 1,
        username: "test",
        music_nums: 0,
      });
      const music_list = ref({
        musics:[]
      });
      const fetchList = async () => {
        // 这里是你的函数逻辑
        // 例如，从服务器获取数据
        // const response = await fetch('http://your-api-endpoint');
        // const data = await response.json();
        // 更新你的响应式数据
        // music_list.musics = data;
        try{
          const response = await axios.get('http://localhost:7986/getlist');
          music_list.value.musics = response.data.musics;
          user.value.music_nums = response.data.count;
        }catch(error){
          console.error('Error fetching data:', error);
        }
      };
      onMounted(()=>{fetchList()});
      return {user,music_list};
    }
  }
</script>

<style scoped>

</style>