<template>
<div @click="MainClick">
    <div v-if="isadmin">
        <h1>Game Page</h1>
        <button @click="AddChallenge">添加题目</button>
    </div>
        <div>
            <h1>{{ this.game.Name }}</h1>
        <div class="divider"></div>
        <div class="Class_" v-if="this.Pwn.length">
            <h1>    PWN     </h1>
            <div class="Challenge" id="pwn"  v-for="pwn in this.Pwn" :key="pwn.ID" @click="ChallengeClick(pwn)">
                <p>{{ pwn.Name }}</p>
                <br>
                <p>{{ pwn.Desc }}</p>
                <br>
                <p>score: {{ pwn.Score }}</p>
            </div>
        </div>
        <div class="divider" v-if="this.Pwn.length"></div>
        <div class="Class_" v-if="this.Reverse.length">
            <h1>    REVERSE     </h1>
            <div class="Challenge" id="reverse"  v-for="re in this.Reverse" :key="re.ID" @click="ChallengeClick(re)">
                <p>{{ re.Name }}</p>
                <br>
                <p>{{ re.Desc }}</p>
                <br>
                <p>score: {{ re.Score }}</p>
            </div>
        </div>
        <div class="divider" v-if="this.Reverse.length"></div>
        <div class="Class_" v-if="this.Web.length">
            <h1>    WEB     </h1>
            <div class="Challenge" id="web"  v-for="web in this.Web" :key="web.ID" @click="ChallengeClick(web)">
                <p>{{ web.Name }}</p>
                <br>
                <p>{{ web.Desc }}</p>
                <br>                
                <p>score: {{ web.Score }}</p>
            </div>
        </div>
        <div class="divider" v-if="this.Web.length"></div>
        <div class="Class_" v-if="this.Crypto.length">
            <h1>    CRYPTO     </h1>
            <div class="Challenge" id="crypto" v-for="cry in this.Crypto" :key="cry.ID" @click="ChallengeClick(cry)">
                <p>{{ cry.Name }}</p>
                <br>
                <p>{{ cry.Desc }}</p>
                <br>
                <p>score: {{ cry.Score }}</p>
            </div>
        </div>
        <div class="divider" v-if="this.Crypto.length"></div>
        <div class="Class_" v-if="this.Misc.length">
            <h1>    MISC     </h1>
            <div class="Challenge" id="misc" v-for="misc in this.Misc" :key="misc.ID" @click="ChallengeClick(misc)">
                <p>{{ misc.Name }}</p>
                <br>
                <p>{{ misc.Desc }}</p>
                <br>
                <p>score: {{ misc.Score }}</p>
            </div>
        </div>
        <div class="divider" v-if="this.Misc.length"></div>
    </div>
    <div v-if="showModala" class="modal">
        <div class="modal-content"  @click="KeepModal">
                <span class="close" @click="closeModal">&times;</span>
                <p>{{ this.s.Name }}</p>
                <br>
                <p>{{ this.s.Desc }}</p>
                <br>
                <p>{{ this.s.Score }}</p>
                <br>
                <div v-if="this.s.Active">
                    <p>image</p>
                </div>
                <form @submit.prevent="FlagSubmit">
                <div>
                    <label for="Flag">Flag:</label>
                    <input type="text" v-model="flag" id="flag" required />
                    <button  type="submit">提交</button>
                </div>
                </form>
                <div v-if="this.s.FileName!=''">
                    <p>file</p>
                </div>
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
            game: {},
            challenge: [],
            Pwn: [],
            Reverse: [],
            Crypto: [],
            Web: [],
            Misc: [],
            isadmin: false,
            showModala: false,
            isfirstclickmain: false,
            s: {},
        };
    },
    async created() {
        if(checkAdmin()){
            this.isadmin = true
        }

        this.id = this.$route.params.id
        this.ChallengeInfo()
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
                /*
                console.log(this.Pwn)
                console.log(this.Reverse)
                console.log(this.Crypto)
                console.log(this.Web)
                console.log(this.Misc)
                */
            }catch(error){
                console.error('Fetch error:', error);
            }
        },

        MainClick(){
            if(!this.isfirstclickmain){
                this.showModala=false
            }
            this.isfirstclickmain=false
            console.log("DDD")
        },
        ChallengeClick(ss){
            this.s=ss
            this.showModala=true
            this.isfirstclickmain=true
        },
        closeModal(){
            this.showModala=false
        },
        KeepModal(){
            this.isfirstclickmain=true
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