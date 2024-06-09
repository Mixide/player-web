<template>
  <NavBar/>
  <router-view/>
  <AudioBar/>
</template>
<script>
import NavBar from './components/NavBar';
import AudioBar from './components/AudioBar';
export default{
  name:'App',
  components:{
    NavBar,
    AudioBar,
  }  
}
</script>

<style>

</style>
