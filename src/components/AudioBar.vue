<template>
    <div id="app">
      <!-- 准备一个容器用来存放音乐播放器 -->
      <div id="aplayer"></div>
    </div>
</template>
  
<script>
  import APlayer from "aplayer"; // 引入音乐插件
  import "aplayer/dist/APlayer.min.css"; // 引入音乐插件的样式
  export default {
    name: "AudioBar",
    data() {
      return {
        audio: [ // 歌曲列表
          {
            name: "我的少年", // 歌曲名字
            artist: "npCswag", // 歌曲演唱者
            url: // 歌曲地址（这里用外链地址）
              "http://localhost:7986/assets/Eviraaa.mp3",
            lrc: "", // 歌词
            theme: "rgb(127, 218, 180)", // 播放这首歌曲时的主题色
          },
        ],
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
    mounted() {
      // 初始化播放器
      this.initAudio();
    },
    methods: {
      initAudio() {
        // 创建一个音乐播放器实例，并挂载到DOM上，同时进行相关配置
        const ap = new APlayer({
          container: document.getElementById("aplayer"),
          audio: this.audio, // 音乐信息
          ...this.info, // 其他配置信息
        });
      },
    },
  };
</script>
  
<style lang="less" scoped>
  #app {
    position:absolute;
    width: 100%;
    bottom: 0;
  }
</style>