<template>
  <div class="card">
    <div class="card-body">
      <div class="row">
        <div class="col-3">
          <input type="file" ref="fileInput"@change="uploadFile"style="display: none;"/>
          <img class="img-fluid" :src="user.photo" alt="选择图片" @click="handleImageClick">
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
import { postImage } from '@/api';
import { mapActions } from 'vuex';
import { useStore } from 'vuex';
export default{
  
  name: "UserInfo",
  props:{
    user:{
        type:Object,
        required:true,
    }
  },
  watch: {
    // 监听 musicList prop 的变化
      user: {}
    },
  data() {
    return {
      files: [] // 用于存储选择的文件
    };
  },
  methods: {
    ...mapActions(['updateUserPhoto']),
    handleImageClick() {
      this.$refs.fileInput.click();
    },
    uploadFile(event) {
      this.files = Array.from(event.target.files);
      if (this.files.length > 0) {
        // 创建一个FormData对象，用于发送文件
        const formData = new FormData();
        this.files.forEach(file => {
          formData.append('file', file);
        });

        // 使用axios发送文件
        postImage(formData,this.user.id).then(response => {
          console.log('文件上传成功', response);
          this.updateUserPhoto(response.data)
        }).catch(error => {
          console.error('文件上传失败', error);
        });
      }
    }
  }
}
</script>

<style scoped>
img {
    border-radius: 50%;
    cursor: pointer;
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