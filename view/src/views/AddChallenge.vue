<template>
  <div>
    <h1>AddChallenge Page</h1>
  </div>
  <div>
    <div id="class_">
      <label for="class_">Class:</label>
      <select id="Class_" v-model="selectedClass">
        <option v-for="option in classOptions" :key="option" :value="option">
          {{ option }}
        </option>
      </select>
    </div>
    <div id="name">
      <label for="name">Name:</label>
      <input type="text" id="Name" v-model="name" name="name">
    </div>
    <div id="desc">
      <label for="description">Description:</label>
      <input type="text" id="Desc" v-model="desc" name="desc">
    </div>
    <div id="score">
      <label for="score">Score:</label>
      <input type="number" id="Score" v-model="score" name="score">
    </div>
    <div id="active">
      <label for="active">Active:</label>
      <input type="checkbox" id="Active" v-model="active" name="active" @change="ActiveChaenge">
    </div>
    <div v-if="image_show" id="image">
      <label for="image">Image:</label>
      <input type="text" id="Image" v-model="image" name="image">
    </div>
    <div id="flag">
      <label for="flag">Flag:</label>
      <input type="text" id="Flag" v-model="flag" name="flag" >
    </div>
    <button type="button" @click="addChallenge">提交</button>
  </div>
</template>
  
  <script>
  import { checkAdmin } from '@/logic/check_jwt';
  import { Url } from '@/config/config';
  export default {
    name: 'AddChallenge',
    data() {
      return {
        classOptions: ['PWN', 'REVERSE', 'WEB', 'CRYPTO', 'MISC'], 
        image_show: false,
        image: '',
        id: '',
      }
    },
    async created() {
      if (!checkAdmin()) {
        this.$router.push('/user');
      }
      this.id = this.$route.params.id
    },
    methods: {
      async addChallenge(){
        const formData = new FormData();
        formData.append('name', localStorage.getItem('name'));
        formData.append('token', localStorage.getItem('jwt'));
        
        formData.append('class', this.selectedClass);
        formData.append('challenge_name', this.name);
        formData.append('active', this.active);
        formData.append('max_score', this.score);
        formData.append('image_name', this.image);
        formData.append('flags', this.flag);
        formData.append('desc', this.desc);
        formData.append('game_id', this.id);
        const res = await fetch(Url + '/challenge/addchallenge', {
          method: 'POST',
          body: formData
        });
        const data = await res.json();
        if (data.status != 200) {
          alert(data.message);
          return;
        }
        alert('添加成功');
        console.log(data);

      },
      ActiveChaenge() {
        this.image_show = this.active;
        if (this.active) {
          this.flag = 'flag{[[FLAG]]}';
        }else{
          if(this.flag == 'flag{[[FLAG]]}')
          this.flag = '';
        }
      },
      IsFileChange() {
        this.file_show = this.is_file;
      },
      handleFileUpload(event) {
        this.file = event.target.files[0];
    },
    }
  }
  </script>

<style>

</style>