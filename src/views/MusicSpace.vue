<template>
  <div id="aplayer"></div>
  <Content>
    <div class="row">
        <div class="col-3">
          <UserInfo :user="user" />
          <Upload :id="user.id"/>
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
  import APlayer from "aplayer"; // 引入音乐插件
  import "aplayer/dist/APlayer.min.css"; // 引入音乐插件的样式
  import {onMounted, reactive, ref} from 'vue';
  import { getMusic } from '@/api';
  import { useStore } from 'vuex';
  export default {
    name: 'MusicSpace',
    components:{
      Content,
      UserInfo,
      MusicList,
      Upload,
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
      const store = useStore();
      const user = ref(store.getters.user || JSON.parse(localStorage.getItem('user')));
      console.log(user.value.photo)
      const music_list = ref({
        musics:[]
      });
      let ap = ref(null);
      const fetchList = async () => {
        try{
          const response = await getMusic(user.value.id);
          music_list.value.musics = response.data;
          user.value.music_nums = response.data.length;
          ap.value = ref(new APlayer({
            container: document.getElementById("aplayer"),
            audio: response.data, // 音乐信息
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
  position: fixed;
  bottom: 0;
  width: 100%;
  z-index: 9999;
}
</style>