<template>
	<div>
		<Weapon
			v-for="(val, idx) in weapons"
			v-bind:weapon="val"
			:key="val.id"
			v-on:remove="removeWeapon(val, idx)"
			:ref="'weapon_' + val.id"
			class="px-0"
		></Weapon>
		<div class="card">
			<h5
				class="header"
				data-toggle="collapse"
				v-bind:data-target="'#new_weapon_' + model_id"
				aria-expanded="false"
				ref="newWeapon"
			>
				New Weapon
			</h5>
			<div class="collapse card-body p-2" v-bind:id="'new_weapon_' + model_id">
				<Weapon :weapon="weapon" v-on:add="addWeapon"></Weapon>
			</div>
		</div>
	</div>
</template>

<script>
import Weapon from './weapon.vue'
export default {
	name: 'Weapons',
	props: ['model_id'],
	components: { Weapon, },
	watch: {
		model_id: function(newVal) {
			this.get(newVal)
		},
	},
	created: function() {
		this.get(this.model_id)
	},
	data() {
		return {
			weapons: [],
			weapon: {
				model_id: this.model_id,
				advantages: [],
			},
		}
	},
	methods: {
		get: function(id) {
			this.weapons = []
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/model/${id}/weapon?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.weapons = res.data
				})
		},
		removeWeapon: function(weapon,index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT + `/weapon/${weapon.id}`)
				.then(function(res) {
					console.debug(res)
					if (res.status === 204) {
						this.weapons.splice(index, 1)
					}
				})
		},
		addWeapon: function(weapon) {
			if (weapon.id == null) {
				weapon.id = 0
				weapon.model_id = this.model_id
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + "/weapon/" + weapon.id + "?lang=" + this.$language, weapon)
				.then(function(res) {
					console.debug(res)
					if (res.status === 201) {
						weapon.id = res.data
						this.weapons.push(weapon)
						this.weapon = {
							model_id: this.model_id,
							advantages: [],
						}
						this.$refs.newWeapon.click()
					}
				})
		},
	},
}
</script>

<style scoped></style>
