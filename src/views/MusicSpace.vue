<template>
  <div id="aplayer"></div>
  <Content>
    <div class="row">
        <div class="col-3">
          <UserInfo :user="user" />
          <Upload/>
        </div>
        <div class="col-9">
          <MusicList :list="music_list"  :ap="ap"/>
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
  import APlayer from "aplayer"; // 引入音乐插件
  import "aplayer/dist/APlayer.min.css"; // 引入音乐插件的样式
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
    data() {
      return {
        info: {
          fixed: false, // 不开启吸底模式
          listFolded: true, // 折叠歌曲列表
          autoplay: false, // 开启自动播放
          preload: "none", // 自动预加载歌曲
          loop: "all", // 播放循环模式、all全部循环 one单曲循环 none只播放一次
          order: "list", //  播放模式，list列表播放, random随机播放
        },
      };
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
      let ap = ref(null);
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
          ap.value = ref(new APlayer({
            container: document.getElementById("aplayer"),
            audio: response.data.musics, // 音乐信息
            fixed: false, // 不开启吸底模式
            listFolded: true, // 折叠歌曲列表
            autoplay: false, // 开启自动播放
            preload: "auto", // 自动预加载歌曲
            loop: "all", // 播放循环模式、all全部循环 one单曲循环 none只播放一次
            order: "list", //  播放模式，list列表播放, random随机播放
          }));
        }catch(error){
          console.error('Error fetching data:', error);
        }
      };
      onMounted(()=>{
        fetchList();
      });
      return {user,music_list,ap};
    }
  }
</script>

<style scoped>
#aplayer {
  position:absolute;
  width: 100%;
  bottom: 0;
}
</style>