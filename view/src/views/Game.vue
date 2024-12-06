<template>
<div @click="MainClick">
    <div v-if="isadmin">
        <h1>Game Page</h1>
        <button @click="AddChallenge">添加题目</button>
    </div>
        <div>
            <div><!--首部-->
            <h1>{{ this.game.Name }}</h1>
            <div id="teamhead">
                <h2>
                    {{ this.teamname }}
                </h2>
            </div>
        </div>
        <div class="divider"></div>
        
        <div>
            <div class="Class_" v-for="type in this.List" >
                <div v-if="type.length">
                    <h2>  {{ type[0].Class }}  </h2>
                    <br>
                    <div class="Challenge" v-for="c_ in type" @click="ChallengeClick(c_)" v-if="type.length">
                        <p>{{ c_.Name }}</p>
                        <br>
                        <p>{{ c_.Desc }}</p>
                        <br>
                        <p>
                            score: {{ c_.Score }}
                            <br>
                            <span>
                                <span>解出: {{ c_.DoneNum }}</span>
                            </span>
                        </p>
                    </div>
                    <div class="divider"></div>
                </div>

            </div>
        </div>

    </div>
    <div v-if="showModala" class="modal">
        <div class="modal-content"  @click="KeepModal">
                <span class="close" @click="closeModal">&times;</span>
                <p>{{ this.s.Name }}</p>
                <br>
                <p>{{ this.s.Desc }}</p>
                <br>
                <p>score: {{ this.s.Score }}</p>
                <br>
                <div v-if="this.s.Active">
                    <p>
                        <div v-if="!this.Live[this.s.ID]"><span><button @click="GetCon(this.s)">创建实例</button></span></div>
                        <div v-if="this.Live[this.s.ID]"><span><button @click="DelCon(this.s)">销毁实例</button></span></div>
                        <div v-if="this.Live[this.s.ID]"><span> 连接: {{ this.Live[this.s.ID] }}</span></div>
                        
                    </p>
                    
                </div>
                <div v-if="this.s.FileName!=''">
                    <p>
                        <span>file: </span>
                        <span><a :href="this.s.FileName" download>{{ this.s.FileName }}</a></span>
                    </p>
                </div>
                <br>
                <form @submit.prevent="FlagSubmit(flag)">
                <div>
                    <label for="Flag">Flag:</label>
                    <input type="text" v-model="flag" id="flag" required />
                    <button  type="submit">提交</button>
                </div>
                </form>
                
            </div>
    </div>
</div>
</template>

<script>
import { Url } from '@/config/config';
import { checkAdmin } from '@/logic/check_jwt';
export default {
    name: 'Game',
    data() {
        return {
            id: '',
            teamname: '',
            teamid: '',
            game: {},
            challenge: [],
            Pwn: [],
            Reverse: [],
            Crypto: [],
            Web: [],
            Misc: [],
            List: [],
            ListName: ['PWN','REVERSE','CRYPTO','WEB','MISC'],
            isadmin: false,
            showModala: false,
            isfirstclickmain: false,
            s: {},
            Live: [],
        };
    },
    async created() {
        if(checkAdmin()){
            this.isadmin = true
        }

        this.id = this.$route.params.id
        this.teamid = this.$route.params.teamid
        this.teamname = this.$route.params.teamname
        this.ChallengeInfo()
        this.GetConInfo()
    },
    methods: {
        AddChallenge() {
            this.$router.push(`/AddChallenge/${this.id}` )
        },
        async ChallengeInfo() {
            const formData = new FormData()
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            formData.append('game_id', this.id)
            try {
                const response = await fetch(`${Url}/game/show`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取题目信息失败');
                }
                const data = await response.json();
                this.game = data.data.game
                this.challenge = Object.values(data.data.challenges);
                for(let i=0;i<this.challenge.length;i++){
                    if(this.challenge[i].Class === 'PWN'){
                        this.Pwn.push(this.challenge[i])
                    }else if(this.challenge[i].Class === 'REVERSE'){
                        this.Reverse.push(this.challenge[i])
                    }else if(this.challenge[i].Class === 'CRYPTO'){
                        this.Crypto.push(this.challenge[i])
                    }else if(this.challenge[i].Class === 'WEB'){
                        this.Web.push(this.challenge[i])
                    }else if(this.challenge[i].Class === 'MISC'){
                        this.Misc.push(this.challenge[i])
                    }
                }
                this.List.push(this.Pwn)
                this.List.push(this.Reverse)
                this.List.push(this.Crypto)
                this.List.push(this.Web)
                this.List.push(this.Misc)
            }catch(error){
                console.error('Fetch error:', error);
            }
        },
        MainClick(){
            if(!this.isfirstclickmain){
                this.showModala=false
            }
            this.isfirstclickmain=false
        },
        ChallengeClick(ss){
            this.s=ss
            this.showModala=true
            this.isfirstclickmain=true
            this.flag=""
        },
        closeModal(){
            this.showModala=false
        },
        KeepModal(){
            this.isfirstclickmain=true
        },
        async GetConInfo(){
            const formData = new FormData()
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            formData.append('game_id', this.id)
            formData.append('team_id', this.teamid)
            try{
                const response = await fetch(`${Url}/game/getconinfo`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取连接信息失败');
                }
                const data = await response.json();
                for(let i =0;i<data.data.length;i++){
                    this.Live[data.data[i].ChallengeID]=data.data[i].Nc
                }
                console.log(this.Live)

            }catch{
                console.error('Fetch error:', error);
            }
        },
        async GetCon(c){
            const formData = new FormData()
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            formData.append('challenge_id', c.ID)
            formData.append('team_id', this.teamid)
            formData.append('game_id', this.id)
            
            try{
                const response = await fetch(`${Url}/challenge/getcon`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取连接失败');
                }
                const data = await response.json();
                this.Live[c.ID]=data.host+':'+data.port
                console.log(this.Live)
            }catch{
                console.error('Fetch error:', error);
            }
        },
        async DelCon(c){
            const formData = new FormData()
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            formData.append('challenge_id', c.ID)
            formData.append('team_id', this.teamid)
            formData.append('game_id', this.id)
            
            try{
                const response = await fetch(`${Url}/challenge/delcon`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取连接失败');
                }
                const data = await response.json();
                delete this.Live[c.ID]
                console.log(this.Live)
            }catch{
                console.error('Fetch error:', error);
            }
        },
        FlagSubmit(flag){

        }
    },
}

</script>

<style>
.Class_ {
  margin-bottom: 20px;
}

.Challenge {
  padding: 10px;
  margin: 10px;
  width: 200px;
  height: 200px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.divider {
  height: 2px;
  background-color: #eee;
  margin: 20px 0;
}

.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 50%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.5);
}
.modal-content {
  background-color: #fefefe;
  margin: auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 500px;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

</style>