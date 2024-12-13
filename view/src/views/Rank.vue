<template>
  <div>
    <h1>Rank Page</h1>
    <div class="RankShow">
      <table>
        <thead>
          <tr>
            <th>排名</th>
            <th>团队名称</th>
            <th>总分</th>
            <th v-for="c in challenge_sort" :key="challenge">
              {{ c.Name }}/<span class="subscript">{{ c.Class }}</span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="team in team_list" :key="team.TeamName">
            <td>{{ team.Rank }}</td>
            <td>{{ team.TeamName }}</td>
            <td>{{ team.Score }}</td>
            <td v-for="c in challenge_sort" :key="challenge">
              <ul v-if="team.Challenge[c.ID]">
                {{ c.Score }}
              </ul>
              <ul v-if="!team.Challenge[c.ID]">
                {{ 0 }}
              </ul>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { Url } from '@/config/config';
export default {
  name: 'Rank',
  data() {
    return {
      gameid: '',
      game: {},
      challenge_list: [],
      pwn_list: new Map(),
      reverse_list: new Map(),
      crypto_list: new Map(),
      web_list: new Map(),
      misc_list: new Map(),
      challenge_sort: [],
      team_list: [],
      page: 1,
      ClassName: ['PWN', 'REVERSE', 'CRYPTO', 'WEB', 'MISC'],
    };
  },
  async created() {
    this.gameid = this.$route.params.id
    this.page = this.$route.params.page

    await this.GetChallengeInfo()
    await this.GetAllTeamRank()
    this.showall()
  },
  methods: {
    async GetChallengeInfo() {
      const formData = new FormData()
      formData.append('name', localStorage.getItem('name'))
      formData.append('token', localStorage.getItem('jwt'))
      formData.append('game_id', this.gameid)
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
        this.challenge_list = data.data.challenges
        for (let i = 0; i < this.challenge_list.length; i++) {
          if (this.challenge_list[i].Class == 'PWN') {
            this.pwn_list.set(this.challenge_list[i].ID, this.challenge_list[i])
          } else if (this.challenge_list[i].Class == 'REVERSE') {
            this.reverse_list.set(this.challenge_list[i].ID, this.challenge_list[i])
          } else if (this.challenge_list[i].Class == 'WEB') {
            this.web_list.set(this.challenge_list[i].ID, this.challenge_list[i])
          } else if (this.challenge_list[i].Class == 'CRYPTO') {
            this.crypto_list.set(this.challenge_list[i].ID, this.challenge_list[i])
          } else if (this.challenge_list[i].Class == 'MISC') {
            this.misc_list.set(this.challenge_list[i].ID, this.challenge_list[i])
          }
          this.challenge_sort = [
            ...Array.from(this.pwn_list.values()),
            ...Array.from(this.reverse_list.values()),
            ...Array.from(this.web_list.values()),
            ...Array.from(this.crypto_list.values()),
            ...Array.from(this.misc_list.values())
          ];


        }
      } catch (error) {
        console.error('Fetch error:', error);
      }
    },
    async GetAllTeamRank() {
      const formData = new FormData()
      formData.append('name', localStorage.getItem('name'))
      formData.append('token', localStorage.getItem('jwt'))
      formData.append('game_id', this.gameid)
      formData.append('page', this.page)
      try {
        const response = await fetch(`${Url}/game/getallteamrank`, {
          method: 'POST',
          body: formData
        });
        const res = await response.json()
        this.team_list = res.data;

      } catch (error) {
        console.error('Fetch error:', error);
      }
    },
    showall() {
      console.log(this.page)
      console.log(this.gameid)
      console.log(this.game)
      console.log(this.challenge_list)
      console.log(this.pwn_list)
      console.log(this.reverse_list)
      console.log(this.web_list)
      console.log(this.crypto_list)
      console.log(this.misc_list)
      console.log(this.team_list)
      console.log(this.challenge_sort)
    }
  }
}
</script>

<style scoped>
.RankShow {
  margin-top: 20px;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  border: 1px solid #ddd;
  padding: 8px;
}
th {
  background-color: #f2f2f2;
  text-align: left;
}
.subscript {
  vertical-align: sub;
  font-size: smaller;
}
</style>