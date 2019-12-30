<template>
	<div>
		<h4>{{ model.name || model.title }} abilities</h4>
		<Ability
			v-for="(val, idx) in abilities"
			v-bind:ability_id="val.id"
			v-bind:ability_type="val.type"
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
				v-bind:data-target="'#new_model_ability' + model.id"
				aria-expanded="false"
			>
				New Ability for {{ model.name || model.title }}
			</h5>
			<div class="collapse card-body p-1" :id="'new_model_ability' + model.id">
				<Ability :abilitiesList="abilitiesList" v-on:add="addAbility" v-on:update="$emit('update')"></Ability>
			</div>
		</div>
		<div class="row">
			<span class="col-1"></span>
			<div class="col-11">
				<WeaponAbilities
					v-for="value in weapons"
					:abilitiesList="abilitiesList"
					:weapon="value"
					:key="value.id"
					v-on:update="$emit('update')"
				></WeaponAbilities>
			</div>
		</div>
		<hr />
	</div>
</template>

<script>
import Ability from './ability.vue'
import WeaponAbilities from './weapon_abilities.vue'
export default {
	name: 'ModelAbilities',
	props: ['model', 'abilitiesList'],
	components: { Ability, WeaponAbilities },
	watch: {
		model: function(newVal) {
			this.get(newVal.id)
			this.getWeapons(newVal.id)
		},
	},
	created: function() {
		this.get(this.model.id)
		this.getWeapons(this.model.id)
	},
	data() {
		return {
			abilities: [],
			weapons: [],
		}
	},
	methods: {
		getWeapons: function(modelID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/model/${modelID}/weapon?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.weapons = res.data
				})
		},
		get: function(modelID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/model/${modelID}/ability?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.abilities = res.data
				})
		},
		removeAbility: function(ability, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT + `/model/${this.model.id}/ability/${ability.id}`)
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
						`/model/${this.model.id}/ability/${ability.id}?type=${ability.type | 0}&lang=${this.$language}`
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
