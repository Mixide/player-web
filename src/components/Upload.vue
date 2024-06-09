<template>
    <div class="card upload-music">
        <div class="card-body">
            <div class="mb-3">
              <label for="formFileMultiple" class="form-label">上传音乐</label>
              <input class="form-control" type="file" id="formFileMultiple" multiple @change="handleFileChange">
              <button type="button" class="btn btn-primary btn-sm" @click="uploadFiles">上传</button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Upload',
  data() {
    return {
      files: [] // 用于存储选择的文件
    };
  },
  methods: {
    handleFileChange(event) {
      // 当文件选择改变时，更新文件列表
      this.files = Array.from(event.target.files);
    },
    uploadFiles() {
      if (this.files.length > 0) {
        // 创建一个FormData对象，用于发送文件
        const formData = new FormData();
        this.files.forEach(file => {
          formData.append('file', file);
        });

        // 使用axios发送文件
        axios.post('/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        .then(response => {
          console.log('文件上传成功', response);
          // 重置文件列表
          this.files = [];
        })
        .catch(error => {
          console.error('文件上传失败', error);
        });
      }
    }
  }
};
</script>

<style scoped>
.upload-music {
    margin-top: 20px;
}
button {
    margin-top: 10px;
}
</style>