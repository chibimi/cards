<template>
	<div>
		<h5 class="my-3">{{ weapon.name || weapon.title }} abilities</h5>
		<Ability
			v-for="(val, idx) in abilities"
			v-bind:ability_id="val.id"
			v-bind:ability_type="val.type"
			v-bind:ability_star="val.star"
			:abilitiesList="abilitiesList"
			:key="val.id"
			v-on:remove="removeAbility(val, idx)"
			v-on:update="$emit('update')"
			v-on:add="addAbility"
		></Ability>
		<div class="card">
			<h5
				class="header"
				data-toggle="collapse"
				v-bind:data-target="'#new_weapon_ability' + weapon.id"
				aria-expanded="false"
			>
				New Ability for {{ weapon.name || weapon.title }}
			</h5>
			<div class="collapse card-body p-2" v-bind:id="'new_weapon_ability' + weapon.id">
				<Ability
					:abilitiesList="abilitiesList"
					v-bind:ability_type="0"
					v-bind:ability_star="0"
					v-on:add="addAbility"
					v-on:update="$emit('update')"
				></Ability>
			</div>
		</div>
	</div>
</template>

<script>
import Ability from './ability.vue'
export default {
	name: 'WeaponAbilities',
	props: ['weapon', 'abilitiesList'],
	components: { Ability },
	watch: {
		weapon: function(newVal) {
			this.get(newVal.id)
		},
	},
	created: function() {
		this.get(this.weapon.id)
	},
	data() {
		return {
			abilities: [],
		}
	},
	methods: {
		get: function(weaponID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/weapon/${weaponID}/ability?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.abilities = res.data
				})
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT + `/weapon/${this.weapon.id}/ability/${ability.id}`)
				.then(function(res) {
					console.debug(res)
					if (res.status === 204) {
						this.abilities.splice(index, 1)
					}
				})
		},
		addAbility: function(ability, push) {
			this.$http
				.put(
					process.env.VUE_APP_API_ENDPOINT +
						`/weapon/${this.weapon.id}/ability/${ability.id}?type=${ability.type}&star=${ability.star}&lang=${this.$language}`
				)
				.then(function(res) {
					console.debug(res)
					if (res.status === 201 && push) {
						this.abilities.push(ability)
					}
				})
		},
	},
}
</script>

<style scoped></style>
